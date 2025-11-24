package main

import (
	"log"
	"net/http"
	"os"

	"cloudCrafters/internal/config"
	"cloudCrafters/internal/db"
	"cloudCrafters/internal/router"
	"cloudCrafters/internal/seed"
)

func main() {
	//Load environment variable
	cfg := config.Load()

	//connect to postgres through GORM
	databse := db.Connect(cfg)
	db.Migrate(databse)

	//check if argument 'seed' is passed
	if len(os.Args) > 1 && os.Args[1] == "seed" {
		seed.Run(databse)
		return
	}

	// Initialize router(mux)
	r := router.NewRouter(databse)

	//Start server
	log.Println("Server running on port: ", cfg.AppPort)
	err := http.ListenAndServe(":"+cfg.AppPort, r)
	if err != nil {
		log.Fatalf("could not start server: %v", err)
		os.Exit(1)
	}
}

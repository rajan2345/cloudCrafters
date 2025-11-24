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
	// Load environment variables
	cfg := config.Load()

	// Connect to database
	database := db.Connect(cfg)
	db.Migrate(database)

	// Handle seeding if passed as argument
	if len(os.Args) > 1 && os.Args[1] == "seed" {
		seed.Run(database)
		log.Println("Database seeded successfully")
		return
	}

	// Setup router
	r := router.NewRouter(database)

	// Handle PORT for cloud deployment
	port := cfg.AppPort
	if port == "" {
		port = os.Getenv("PORT")
	}
	if port == "" {
		port = "8080"
	}

	// Add basic health check
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")

	// Start server
	log.Println("Server running on port:", port)
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}

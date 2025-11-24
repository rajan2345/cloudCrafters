package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"cloudCrafters/internal/mappings"
	"cloudCrafters/internal/services"
)

func NewRouter(db *gorm.DB) *mux.Router {
	r := mux.NewRouter()

	//Initialize the repositories
	serviceRepo := services.NewServiceRepository(db)
	mappingRepo := mappings.NewMappingRepository(db)

	//Initialize the handlers
	serviceHandler := services.NewServiceHandler(serviceRepo)
	mappingHandler := mappings.NewMappingHanlder(mappingRepo, serviceRepo)

	// Healthe routers
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	}).Methods("GET")

	//service routes
	r.HandleFunc("/services", serviceHandler.GetAllServices).Methods("GET")
	r.HandleFunc("/services/{provider}", serviceHandler.GetByProvider).Methods("GET")
	r.HandleFunc("/services/{provider}/{code}", serviceHandler.GetByCode).Methods("GET")

	//Mapping route
	r.HandleFunc("/mapping", mappingHandler.GetMapping).Methods("GET")

	return r
}

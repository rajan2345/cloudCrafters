package services

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type ServiceHandler struct {
	Repo *ServiceRepository
}

// Constructor: used in router
func NewServiceHandler(repo *ServiceRepository) *ServiceHandler {
	return &ServiceHandler{Repo: repo}
}

// GET Services
func (h *ServiceHandler) GetAllServices(w http.ResponseWriter, r *http.Request) {
	services, err := h.Repo.GetAll()
	if err != nil {
		http.Error(w, "failed to fetch services", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(services)
}

// GET /services/(provider)
func (h *ServiceHandler) GetByProvider(w http.ResponseWriter, r *http.Request) {
	provider := mux.Vars(r)["provider"]

	if provider == "" {
		http.Error(w, "provider is required", http.StatusBadRequest)
		return
	}

	services, err := h.Repo.GetByProvider(provider)

	if err != nil {
		http.Error(w, "failed to fethc services", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(services)
}

// GET /services/{provider}/{code}
func (h *ServiceHandler) GetByCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	provider := vars["provider"]
	code := vars["code"]

	if provider == "" || code == "" {
		http.Error(w, "provider and code are required", http.StatusBadRequest)
		return
	}

	service, err := h.Repo.Get(provider, code)
	if err != nil {
		http.Error(w, "service not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(service)
}

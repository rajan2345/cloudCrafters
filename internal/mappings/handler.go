package mappings

import (
	"cloudCrafters/internal/services"
	"encoding/json"
	"net/http"
)

// MappingHandler connects mapping repository + service repository
type MappingHandler struct {
	MappingRepo *MappingRepository
	ServiceRepo *services.ServiceRepository
}

// Constructor for initialize MappingHandler
func NewMappingHanlder(mappingRepo *MappingRepository, serviceRepo *services.ServiceRepository) *MappingHandler {
	return &MappingHandler{
		MappingRepo: mappingRepo,
		ServiceRepo: serviceRepo,
	}
}

// GET /mapping?from=aws&service=ec2&to=gcp -- e.g
func (h *MappingHandler) GetMapping(w http.ResponseWriter, r *http.Request) {
	//Read Query params
	fromProvider := r.URL.Query().Get("from")
	serviceCode := r.URL.Query().Get("service")
	toProvider := r.URL.Query().Get("to")

	//Validate required fields
	if fromProvider == "" || serviceCode == "" || toProvider == "" {
		http.Error(w, "query  params 'from', 'service' and 'to' is required", http.StatusBadRequest)
		return
	}

	//Fetching Map repository
	mapping, err := h.MappingRepo.GetMapping(fromProvider, serviceCode, toProvider)
	if err != nil {
		http.Error(w, "no mapping found", http.StatusNotFound)
		return
	}

	// Fetch full details from mapping variable
	targetService, err := h.ServiceRepo.Get(mapping.FromProvider, mapping.ToCode)
	if err != nil {
		http.Error(w, "target service not found", http.StatusInternalServerError)
		return
	}

	// Build final response from the API
	response := map[string]interface{}{
		"from": map[string]string{
			"provider": fromProvider,
			"service":  serviceCode,
		},
		"to": map[string]interface{}{
			"provider": targetService.Provider,
			"service":  targetService.Code,
			"name":     targetService.Name,
			"category": targetService.Category,
		},
	}

	//Return json
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

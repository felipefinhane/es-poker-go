package participant

import (
	"encoding/json"
	"net/http"

	"github.com/felipefinhane/es-poker-go/internal/app/participant"
)

type ParticipantHandler struct {
	service *participant.ParticipantService
}

func NewParticipantHandler(service *participant.ParticipantService) *ParticipantHandler {
	return &ParticipantHandler{service: service}
}

func (h *ParticipantHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	participant, err := h.service.RegisterParticipant(r.Context(), req.Name, req.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(participant)
}

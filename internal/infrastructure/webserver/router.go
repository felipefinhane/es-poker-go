package webserver

import (
	"github.com/felipefinhane/es-poker-go/internal/app/participant"
	"github.com/gorilla/mux"
)

func NewRouter(participantHandler *participant.ParticipantHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/participants/register", participantHandler.Register).Methods("POST")
	return r
}

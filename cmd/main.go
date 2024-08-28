package main

import (
	"log"
	"net/http"

	"github.com/felipefinhane/es-poker-go/internal/app/event"
	"github.com/felipefinhane/es-poker-go/internal/app/participant"
	"github.com/felipefinhane/es-poker-go/internal/infrastructure/persistence"
	"github.com/felipefinhane/es-poker-go/internal/infrastructure/persistence/mongo"
	"github.com/felipefinhane/es-poker-go/internal/infrastructure/webserver"
)

func main() {
	db, err := persistence.ConnectDB("mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}

	eventRepo := mongo.NewEventRepository(db)
	eventService := event.NewEventService(eventRepo)

	participantRepo := mongo.NewParticipantRepository(db)
	participantService := participant.NewParticipantService(participantRepo, *eventService)
	participantHandler := participant.NewParticipantHandler(participantService)

	r := webserver.NewRouter(participantHandler)

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

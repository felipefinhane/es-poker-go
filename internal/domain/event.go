package domain

import "time"

type Event struct {
	EntityID   string
	EntityType string
	EventType  string
	EventData  map[string]interface{}
	OccurredAt time.Time
}

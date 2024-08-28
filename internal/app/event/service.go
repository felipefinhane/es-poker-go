package event

import (
	"context"
	"time"

	"github.com/felipefinhane/es-poker-go/internal/domain"
)

type EventService struct {
	repo EventRepository
}

type EventRepository interface {
	StoreEvent(ctx context.Context, event domain.Event) error
	GetEventsByEntity(ctx context.Context, entityID, entityType string) ([]domain.Event, error)
}

func NewEventService(repo EventRepository) *EventService {
	return &EventService{repo: repo}
}

func (s *EventService) CreateEvent(ctx context.Context, entityID, entityType, eventType string, eventData map[string]interface{}) error {
	event := domain.Event{
		EntityID:   entityID,
		EntityType: entityType,
		EventType:  eventType,
		EventData:  eventData,
		OccurredAt: time.Now(),
	}
	return s.repo.StoreEvent(ctx, event)
}

func (s *EventService) GetEventsByEntity(ctx context.Context, entityID, entityType string) ([]domain.Event, error) {
	return s.repo.GetEventsByEntity(ctx, entityID, entityType)
}

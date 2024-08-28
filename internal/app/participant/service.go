package participant

import (
	"context"
	"time"

	"github.com/felipefinhane/es-poker-go/internal/app/event"
	"github.com/felipefinhane/es-poker-go/internal/domain"
	"github.com/google/uuid"
)

type ParticipantService struct {
	repo   ParticipantRepository
	events event.EventService
}

type ParticipantRepository interface {
	Save(ctx context.Context, participant domain.Participant) error
	FindByID(ctx context.Context, id string) (*domain.Participant, error)
}

func NewParticipantService(repo ParticipantRepository, events event.EventService) *ParticipantService {
	return &ParticipantService{
		repo:   repo,
		events: events,
	}
}

func (s *ParticipantService) RegisterParticipant(ctx context.Context, name, email string) (*domain.Participant, error) {
	id := uuid.New().String()
	participant := domain.Participant{
		ID:       id,
		Name:     name,
		Email:    email,
		JoinedAt: time.Now(),
	}

	if err := s.repo.Save(ctx, participant); err != nil {
		return nil, err
	}

	eventData := map[string]interface{}{
		"name":  name,
		"email": email,
	}
	if err := s.events.CreateEvent(ctx, id, "Participant", "ParticipantRegistered", eventData); err != nil {
		return nil, err
	}

	return &participant, nil
}

func (s *ParticipantService) GetParticipantByID(ctx context.Context, id string) (*domain.Participant, error) {
	return s.repo.FindByID(ctx, id)
}

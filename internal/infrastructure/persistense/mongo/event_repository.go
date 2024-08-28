package mongo

import (
	"context"

	"github.com/felipefinhane/es-poker-go/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type EventRepository struct {
	collection *mongo.Collection
}

func NewEventRepository(db *mongo.Database) *EventRepository {
	return &EventRepository{
		collection: db.Collection("events"),
	}
}

func (r *EventRepository) StoreEvent(ctx context.Context, event domain.Event) error {
	_, err := r.collection.InsertOne(ctx, event)
	return err
}

func (r *EventRepository) GetEventsByEntity(ctx context.Context, entityID, entityType string) ([]domain.Event, error) {
	var events []domain.Event
	cursor, err := r.collection.Find(ctx, bson.M{
		"entityid":   entityID,
		"entitytype": entityType,
	})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var event domain.Event
		if err := cursor.Decode(&event); err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

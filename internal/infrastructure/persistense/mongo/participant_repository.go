package mongo

import (
	"context"

	"github.com/felipefinhane/es-poker-go/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type ParticipantRepository struct {
	collection *mongo.Collection
}

func NewParticipantRepository(db *mongo.Database) *ParticipantRepository {
	return &ParticipantRepository{
		collection: db.Collection("participants"),
	}
}

func (r *ParticipantRepository) Save(ctx context.Context, participant domain.Participant) error {
	_, err := r.collection.InsertOne(ctx, participant)
	return err
}

func (r *ParticipantRepository) FindByID(ctx context.Context, id string) (*domain.Participant, error) {
	var participant domain.Participant
	err := r.collection.FindOne(ctx, bson.M{"id": id}).Decode(&participant)
	if err != nil {
		return nil, err
	}
	return &participant, nil
}

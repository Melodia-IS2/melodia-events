package persistence

import (
	"context"
	"melodia-events/internal/domain/entities"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MongoEventRepository struct {
	Collection *mongo.Collection
}

func (r *MongoEventRepository) Register(ctx context.Context, event *entities.Event) error {
	_, err := r.Collection.InsertOne(ctx, event)
	return err
}

func (r *MongoEventRepository) FindAll(ctx context.Context) ([]*entities.Event, error) {
	cursor, err := r.Collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	events := []*entities.Event{}
	for cursor.Next(ctx) {
		var event entities.Event
		if err := cursor.Decode(&event); err != nil {
			return nil, err
		}
		events = append(events, &event)
	}
	return events, nil
}

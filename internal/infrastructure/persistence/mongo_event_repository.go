package persistence

import (
	"context"
	"melodia-events/internal/domain/entities"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
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

func (r *MongoEventRepository) FindDueUnpublished(ctx context.Context, before time.Time) ([]*entities.Event, error) {
	filter := bson.D{
		{Key: "publish.is_published", Value: false},
		{Key: "publish.publish_after", Value: bson.D{{Key: "$lte", Value: before}}},
	}

	findOpts := options.Find().SetSort(bson.D{{Key: "publish.publish_after", Value: 1}})

	cursor, err := r.Collection.Find(ctx, filter, findOpts)
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

func (r *MongoEventRepository) MarkPublished(ctx context.Context, id uuid.UUID) error {
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "publish.is_published", Value: true}}}}
	_, err := r.Collection.UpdateOne(ctx, filter, update)
	return err
}

package persistence

import (
	"context"

	"github.com/Melodia-IS2/melodia-go-utils/pkg/logger"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MongoLogRepository struct {
	Collection *mongo.Collection
}

func (r *MongoLogRepository) Log(ctx context.Context, log *logger.Log) error {
	_, err := r.Collection.InsertOne(ctx, log)
	return err
}

func (r *MongoLogRepository) FindAll(ctx context.Context) ([]*logger.Log, error) {
	cursor, err := r.Collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	logs := []*logger.Log{}
	for cursor.Next(ctx) {
		var log logger.Log
		if err := cursor.Decode(&log); err != nil {
			return nil, err
		}
		logs = append(logs, &log)
	}
	return logs, nil
}

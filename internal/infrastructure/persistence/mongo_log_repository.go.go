package persistence

import (
	"context"

	"github.com/Melodia-IS2/melodia-events/internal/domain/entities"
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
func (r *MongoLogRepository) Search(ctx context.Context, search entities.LogSearch) ([]*logger.Log, error) {
	filter := bson.M{}

	if search.DateFrom != nil || search.DateTo != nil {
		dateFilter := bson.M{}
		if search.DateFrom != nil {
			dateFilter["$gte"] = search.DateFrom
		}
		if search.DateTo != nil {
			dateFilter["$lte"] = search.DateTo
		}
		filter["created_at"] = dateFilter
	}

	if search.OnlyEntries != nil && *search.OnlyEntries {
		filter["entries.0"] = bson.M{"$exists": true}
	}

	if search.Level != nil {
		filter["level"] = *search.Level
	}

	if search.Application != nil {
		filter["application"] = *search.Application
	}

	if search.Endpoint != nil {
		filter["endpoint"] = *search.Endpoint
	}

	if search.Method != nil {
		filter["method"] = *search.Method
	}

	if search.Status != nil {
		filter["status"] = *search.Status
	}

	if search.EntriesMessage != nil {
		filter["entries.message"] = *search.EntriesMessage
	}

	if search.EntriesLevel != nil {
		filter["entries.level"] = *search.EntriesLevel
	}

	cursor, err := r.Collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var logs []*logger.Log
	for cursor.Next(ctx) {
		var log logger.Log
		if err := cursor.Decode(&log); err != nil {
			return nil, err
		}
		logs = append(logs, &log)
	}

	return logs, nil
}

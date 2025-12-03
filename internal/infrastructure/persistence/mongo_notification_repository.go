package persistence

import (
	"context"

	"github.com/Melodia-IS2/melodia-events/internal/domain/entities"
	"github.com/google/uuid"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoNotificationsRepository struct {
	Collection *mongo.Collection
}

func (r *MongoNotificationsRepository) Register(ctx context.Context, notification *entities.Notification) error {
	_, err := r.Collection.InsertOne(ctx, notification)
	return err
}

func (r *MongoNotificationsRepository) Get(
	ctx context.Context,
	n uint,
	userID uuid.UUID,
) ([]*entities.Notification, error) {

	filter := bson.M{"user_id": userID}

	opts := options.Find().
		SetSort(bson.D{{"created_at", -1}}).
		SetLimit(int64(n))

	cursor, err := r.Collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	notifications := []*entities.Notification{}
	for cursor.Next(ctx) {
		var notification entities.Notification
		if err := cursor.Decode(&notification); err != nil {
			return nil, err
		}
		notifications = append(notifications, &notification)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return notifications, nil
}

func (r *MongoNotificationsRepository) MarkAsRead(ctx context.Context, notificationID uuid.UUID, userID uuid.UUID) error {
	_, err := r.Collection.UpdateOne(ctx, bson.M{"_id": notificationID, "user_id": userID}, bson.M{"$set": bson.M{"read": true}})
	return err
}

package entities

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	ID        uuid.UUID         `json:"id" bson:"_id"`
	UserID    uuid.UUID         `json:"user_id" bson:"user_id"`
	Read      bool              `json:"read" bson:"read"`
	Data      map[string]string `json:"data" bson:"data"`
	CreatedAt time.Time         `json:"created_at" bson:"created_at"`
}

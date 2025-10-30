package entities

import (
	"time"

	"github.com/google/uuid"
)

type EventStatus int

const (
	EventStatusPending EventStatus = iota
	EventStatusProcessing
	EventStatusPublished
	EventStatusFailed
)

type Event struct {
	ID           uuid.UUID      `json:"id" bson:"_id"`
	Topic        string         `json:"topic" bson:"topic"`
	Title        string         `json:"title" bson:"title"`
	Payload      map[string]any `json:"payload" bson:"payload"`
	PublishAfter time.Time      `json:"publish_after" bson:"publish_after"`
	CreatedAt    time.Time      `json:"created_at" bson:"created_at"`
}

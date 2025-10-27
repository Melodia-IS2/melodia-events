package entities

import (
	"time"

	"github.com/google/uuid"
)

type LogLevel int

const (
	DebugLevel   LogLevel = 0
	InfoLevel    LogLevel = 1
	WarningLevel LogLevel = 2
	ErrorLevel   LogLevel = 3
)

type Event struct {
	ID        uuid.UUID     `json:"id" bson:"_id"`
	Publish   *PublishEvent `json:"publish" bson:"publish"`
	Log       *LogEvent     `json:"log" bson:"log"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	Service   string        `json:"service" bson:"service"`
	Token     string        `json:"token" bson:"token"`
}

type LogEvent struct {
	Code    string         `json:"code" bson:"code"`
	Message string         `json:"message" bson:"message"`
	Payload map[string]any `json:"payload" bson:"payload"`
	Level   LogLevel       `json:"level" bson:"level"`
}

type PublishEvent struct {
	Topic        string         `json:"topic" bson:"topic"`
	Title        string         `json:"title" bson:"title"`
	Payload      map[string]any `json:"payload" bson:"payload"`
	PublishAfter time.Time      `json:"publish_after" bson:"publish_after"`
	IsPublished  bool           `json:"is_published" bson:"is_published"`
}

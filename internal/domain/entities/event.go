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
	ID        uuid.UUID     `json:"id"`
	Publish   *PublishEvent `json:"publish"`
	Log       *LogEvent     `json:"log"`
	CreatedAt time.Time     `json:"created_at"`
	Service   string        `json:"service"`
	Token     string        `json:"token"`
}

type LogEvent struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Payload map[string]any `json:"payload"`
	Level   LogLevel       `json:"level"`
}

type PublishEvent struct {
	Topic        string         `json:"topic"`
	Title        string         `json:"title"`
	Payload      map[string]any `json:"payload"`
	PublishAfter time.Time      `json:"publish_after"`
	IsPublished  bool           `json:"is_published"`
}

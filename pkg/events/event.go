package events

import (
	"time"

	"github.com/Melodia-IS2/melodia-events/internal/domain/entities"
	"github.com/google/uuid"
)

type Event struct {
	Topic        string         `json:"topic"`
	Key          string         `json:"key"`
	Payload      map[string]any `json:"payload"`
	PublishAfter time.Time      `json:"publish_after"`
}

func (e *Event) ToDomain() *entities.Event {
	return &entities.Event{
		ID:           uuid.New(),
		Topic:        e.Topic,
		Key:          e.Key,
		Payload:      e.Payload,
		PublishAfter: e.PublishAfter,
		CreatedAt:    time.Now(),
	}
}

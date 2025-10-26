package events

import (
	"context"
	"melodia-events/internal/domain/entities"
	"time"
)

var service string

func SetService(s string) {
	service = s
}

type EventBuilder struct {
	event *entities.Event
}

func New(ctx context.Context) *EventBuilder {
	token := ctx.Value("token").(string)

	return &EventBuilder{
		event: &entities.Event{
			CreatedAt: time.Now(),
			Service:   service,
			Token:     token,
		},
	}
}

func (b *EventBuilder) Log(code string, message string, level entities.LogLevel, payload *map[string]any) *EventBuilder {
	if payload == nil {
		payload = &map[string]any{}
	}
	b.event.Log = &entities.LogEvent{
		Code:    code,
		Message: message,
		Level:   level,
		Payload: *payload,
	}
	return b
}

func (b *EventBuilder) Publish(topic string, title string, payload *map[string]any) *EventBuilder {
	if payload == nil {
		payload = &map[string]any{}
	}
	b.event.Publish = &entities.PublishEvent{
		Topic:        topic,
		Title:        title,
		Payload:      *payload,
		PublishAfter: time.Now(),
		IsPublished:  false,
	}
	return b
}
func (b *EventBuilder) ProgrammedPublish(topic string, title string, payload *map[string]any, publishAfter time.Time) *EventBuilder {
	if payload == nil {
		payload = &map[string]any{}
	}
	b.event.Publish = &entities.PublishEvent{
		Topic:        topic,
		Title:        title,
		Payload:      *payload,
		PublishAfter: publishAfter,
		IsPublished:  false,
	}
	return b
}

func (b *EventBuilder) Build() *entities.Event {
	return b.event
}

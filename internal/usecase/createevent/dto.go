package createevent

import (
	"melodia-events/internal/domain/entities"
)

type CreateEventRequest struct {
	Topic       string `form:"topic"`
	Title       string `form:"title"`
	Description string `form:"description"`
	Service     string `form:"service"`
	Payload     string `form:"payload"`
	ErrorLevel  int    `form:"error_level"`
	Token       string `form:"token"`
}

func (r *CreateEventRequest) ToDomain() (*entities.Event, error) {
	return &entities.Event{
		Topic:       r.Topic,
		Title:       r.Title,
		Description: r.Description,
		Service:     r.Service,
		Payload:     r.Payload,
		ErrorLevel:  entities.ErrorLevel(r.ErrorLevel),
		Token:       r.Token,
	}, nil
}

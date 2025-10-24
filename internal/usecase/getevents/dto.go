package getevents

import (
	"melodia-events/internal/domain/entities"
)

type GetEventsResponse struct {
	ID          string `json:"id"`
	Topic       string `json:"topic"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Service     string `json:"service"`
	Payload     string `json:"payload"`
	ErrorLevel  int    `json:"error_level"`
	Token       string `json:"token"`
}

func FromDomain(event *entities.Event) GetEventsResponse {
	return GetEventsResponse{
		ID:          event.ID.String(),
		Topic:       event.Topic,
		Title:       event.Title,
		Description: event.Description,
		Service:     event.Service,
		Payload:     event.Payload,
		ErrorLevel:  int(event.ErrorLevel),
		Token:       event.Token,
	}
}

func FromDomainArray(events []*entities.Event) []GetEventsResponse {
	requests := make([]GetEventsResponse, len(events))
	for _, event := range events {
		requests = append(requests, FromDomain(event))
	}
	return requests
}

package getevents

import (
	"context"

	"melodia-events/internal/domain/entities"
	"melodia-events/internal/domain/repositories"
)

type GetEvents interface {
	Execute(ctx context.Context) ([]*entities.Event, error)
}

type GetEventsImpl struct {
	EventRepository repositories.EventRepository
}

func (u *GetEventsImpl) Execute(ctx context.Context) ([]*entities.Event, error) {
	return u.EventRepository.FindAll(ctx)
}

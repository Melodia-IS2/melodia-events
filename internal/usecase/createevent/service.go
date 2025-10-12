package createevent

import (
	"context"

	"melodia-events/internal/domain/entities"
	"melodia-events/internal/domain/repositories"
)

type CreateEvent interface {
	Execute(ctx context.Context, event *entities.Event) error
}

type CreateEventImpl struct {
	EventRepository repositories.EventRepository
	EventPublisher  repositories.EventPublisher
}

func (u *CreateEventImpl) Execute(ctx context.Context, event *entities.Event) (err error) {

	err = u.EventRepository.Register(ctx, event)
	if err != nil {
		return err
	}

	if err := u.EventPublisher.Publish(ctx, event); err != nil {
		return err
	}

	return nil
}

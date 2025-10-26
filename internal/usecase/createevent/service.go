package createevent

import (
	"context"

	"melodia-events/internal/domain/entities"
	"melodia-events/internal/domain/repositories"

	"github.com/google/uuid"
)

type CreateEvent interface {
	Execute(ctx context.Context, event *entities.Event) error
}

type CreateEventImpl struct {
	EventRepository repositories.EventRepository
	EventPublisher  repositories.EventPublisher
}

func (u *CreateEventImpl) Execute(ctx context.Context, event *entities.Event) (err error) {
	event.ID = uuid.New()

	if event.Publish != nil {
		if err := u.EventPublisher.Publish(ctx, event); err != nil {
			print(err.Error())
			return err
		}
	}

	if event.Log != nil {
		if err := u.EventRepository.Register(ctx, event); err != nil {
			print(err.Error())
			return err
		}
	}

	return nil
}

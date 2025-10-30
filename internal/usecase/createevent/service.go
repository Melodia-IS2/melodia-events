package createevent

import (
	"context"
	"time"

	"github.com/Melodia-IS2/melodia-events/internal/domain/entities"
	"github.com/Melodia-IS2/melodia-events/internal/domain/repositories"

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

	if event.PublishAfter.Before(time.Now()) {
		if err := u.EventPublisher.Publish(ctx, event); err != nil {
			print(err.Error())
			return err
		}
		return nil
	}

	if err := u.EventRepository.Schedule(ctx, *event); err != nil {
		print(err.Error())
		return err
	}

	return nil
}

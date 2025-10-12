package repositories

import (
	"context"
	"melodia-events/internal/domain/entities"
)

type EventPublisher interface {
	Publish(ctx context.Context, event *entities.Event) error
}

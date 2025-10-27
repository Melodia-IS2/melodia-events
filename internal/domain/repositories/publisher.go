package repositories

import (
	"context"

	"github.com/Melodia-IS2/melodia-events/internal/domain/entities"
)

type EventPublisher interface {
	Publish(ctx context.Context, event *entities.Event) error
}

package repositories

import (
	"context"
	"melodia-events/internal/domain/entities"
)

type EventRepository interface {
	Register(ctx context.Context, event *entities.Event) error
}

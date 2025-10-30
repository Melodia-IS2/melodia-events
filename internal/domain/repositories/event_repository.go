package repositories

import (
	"context"

	"github.com/Melodia-IS2/melodia-events/internal/domain/entities"
)

type EventRepository interface {
	Schedule(ctx context.Context, e entities.Event) error
	FetchDueEvents(ctx context.Context, limit int64) ([]entities.Event, error)
	FindAll(ctx context.Context) ([]entities.Event, error)
}

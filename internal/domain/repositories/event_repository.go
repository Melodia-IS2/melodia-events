package repositories

import (
	"context"
	"melodia-events/internal/domain/entities"
	"time"

	"github.com/google/uuid"
)

type EventRepository interface {
	Register(ctx context.Context, event *entities.Event) error
	FindAll(ctx context.Context) ([]*entities.Event, error)
	FindDueUnpublished(ctx context.Context, before time.Time) ([]*entities.Event, error)
	MarkPublished(ctx context.Context, id uuid.UUID) error
}

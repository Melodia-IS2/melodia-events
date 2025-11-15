package repositories

import (
	"context"

	"github.com/Melodia-IS2/melodia-events/internal/domain/entities"
	"github.com/google/uuid"
)

type DevicesRepository interface {
	Register(ctx context.Context, device entities.Device) error
	Delete(ctx context.Context, device entities.Device) error
	FetchByUserIDs(ctx context.Context, ids []uuid.UUID) ([]entities.Device, error)
}

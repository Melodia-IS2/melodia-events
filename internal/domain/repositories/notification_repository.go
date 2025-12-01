package repositories

import (
	"context"

	"github.com/Melodia-IS2/melodia-events/internal/domain/entities"
	"github.com/google/uuid"
)

type NotificationsRepository interface {
	Register(ctx context.Context, notification *entities.Notification) error
	Get(ctx context.Context, n uint, userID uuid.UUID) ([]*entities.Notification, error)
	MarkAsRead(ctx context.Context, notificationID uuid.UUID) error
}

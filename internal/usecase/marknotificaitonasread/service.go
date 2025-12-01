package marknotificaitonasread

import (
	"context"

	"github.com/Melodia-IS2/melodia-events/internal/domain/repositories"
	"github.com/google/uuid"
)

type MarkNotificationAsRead interface {
	Execute(ctx context.Context, notificationID uuid.UUID) error
}

type MarkNotificationAsReadImpl struct {
	NotificationsRepository repositories.NotificationsRepository
}

func (u *MarkNotificationAsReadImpl) Execute(ctx context.Context, notificationID uuid.UUID) error {
	return u.NotificationsRepository.MarkAsRead(ctx, notificationID)
}

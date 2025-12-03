package getnotifications

import (
	"context"

	"github.com/Melodia-IS2/melodia-events/internal/domain/entities"
	"github.com/Melodia-IS2/melodia-events/internal/domain/repositories"
	"github.com/google/uuid"
)

type GetNotifications interface {
	Execute(ctx context.Context, n int, userID uuid.UUID) ([]*entities.Notification, error)
}

type GetNotificationsImpl struct {
	NotificationsRepository repositories.NotificationsRepository
}

func (u *GetNotificationsImpl) Execute(ctx context.Context, n int, userID uuid.UUID) ([]*entities.Notification, error) {
	return u.NotificationsRepository.Get(ctx, n, userID)
}

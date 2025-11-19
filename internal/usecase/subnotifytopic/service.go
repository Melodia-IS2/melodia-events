package subnotifytopic

import (
	"context"

	firebase "firebase.google.com/go/v4"

	"github.com/Melodia-IS2/melodia-events/internal/domain/repositories"
	"github.com/google/uuid"
)

type SubNotifyTopic interface {
	SubNotifyTopic(ctx context.Context, userID uuid.UUID, topic string) error
	UnsubNotifyTopic(ctx context.Context, userID uuid.UUID, topic string) error
}

type SubNotifyTopicImpl struct {
	FirebaseApp       *firebase.App
	DevicesRepository repositories.DevicesRepository
}

func (u *SubNotifyTopicImpl) SubNotifyTopic(ctx context.Context, userID uuid.UUID, topic string) (err error) {
	client, err := u.FirebaseApp.Messaging(ctx)
	if err != nil {
		return err
	}

	devices, err := u.DevicesRepository.FetchByUserIDs(ctx, []uuid.UUID{userID})
	if err != nil {
		return err
	}

	deviceTokens := make([]string, 0, len(devices))
	for _, device := range devices {
		deviceTokens = append(deviceTokens, device.DeviceToken)
	}

	client.SubscribeToTopic(ctx, deviceTokens, topic)

	return nil
}

func (u *SubNotifyTopicImpl) UnsubNotifyTopic(ctx context.Context, userID uuid.UUID, topic string) (err error) {
	client, err := u.FirebaseApp.Messaging(ctx)
	if err != nil {
		return err
	}

	devices, err := u.DevicesRepository.FetchByUserIDs(ctx, []uuid.UUID{userID})
	if err != nil {
		return err
	}

	deviceTokens := make([]string, 0, len(devices))
	for _, device := range devices {
		deviceTokens = append(deviceTokens, device.DeviceToken)
	}

	client.UnsubscribeFromTopic(ctx, deviceTokens, topic)

	return nil
}

package notify

import (
	"context"
	"fmt"
	"time"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"

	"github.com/Melodia-IS2/melodia-events/internal/domain/entities"
	"github.com/Melodia-IS2/melodia-events/internal/domain/repositories"
	"github.com/google/uuid"
)

type Notify interface {
	NotifyUser(ctx context.Context, userID uuid.UUID, key string, data map[string]string) error
	NotifyUsers(ctx context.Context, userIDs []uuid.UUID, key string, data map[string]string) error
	NotifyTopic(ctx context.Context, topic string, key string, data map[string]string) error
}

type NotifyImpl struct {
	FirebaseApp             *firebase.App
	DevicesRepository       repositories.DevicesRepository
	NotificationsRepository repositories.NotificationsRepository
}

func (u *NotifyImpl) NotifyUser(ctx context.Context, userID uuid.UUID, key string, data map[string]string) (err error) {

	client, err := u.FirebaseApp.Messaging(ctx)
	if err != nil {
		return err
	}

	devices, err := u.DevicesRepository.FetchByUserIDs(ctx, []uuid.UUID{userID})
	if err != nil {
		return err
	}

	userTokens := make([]string, 0, len(devices))
	for _, device := range devices {
		userTokens = append(userTokens, device.DeviceToken)
	}

	err = u.NotificationsRepository.Register(ctx, &entities.Notification{
		ID:        uuid.New(),
		UserID:    userID,
		Topic:     key,
		Data:      data,
		Read:      false,
		CreatedAt: time.Now(),
	})
	if err != nil {
		return err
	}

	client.SendEachForMulticast(ctx, &messaging.MulticastMessage{
		Tokens: userTokens,
		Data:   data,
	})
	return nil
}

func (u *NotifyImpl) NotifyTopic(ctx context.Context, topic string, key string, data map[string]string) (err error) {

	client, err := u.FirebaseApp.Messaging(ctx)
	if err != nil {
		return err
	}

	client.Send(ctx, &messaging.Message{
		Topic: topic,
		Data:  data,
	})
	return nil
}

func (u *NotifyImpl) NotifyUsers(ctx context.Context, userIDs []uuid.UUID, key string, data map[string]string) (err error) {

	fmt.Println("NotifyUsers", userIDs, key, data)
	client, err := u.FirebaseApp.Messaging(ctx)
	if err != nil {
		return err
	}

	fmt.Println("Fetching devices")
	devices, err := u.DevicesRepository.FetchByUserIDs(ctx, userIDs)
	if err != nil {
		return err
	}

	userTokens := make([]string, 0, len(devices))
	for _, device := range devices {
		userTokens = append(userTokens, device.DeviceToken)
	}

	for _, userID := range userIDs {
		err = u.NotificationsRepository.Register(ctx, &entities.Notification{
			ID:        uuid.New(),
			UserID:    userID,
			Topic:     key,
			Data:      data,
			Read:      false,
			CreatedAt: time.Now(),
		})
		if err != nil {
			fmt.Println("Error registering notification", err.Error())
			continue
		}
	}

	fmt.Println("Sending notifications to", userTokens)
	_, err = client.SendEachForMulticast(ctx, &messaging.MulticastMessage{
		Tokens: userTokens,
		Data:   data,
	})

	if err != nil {
		fmt.Println("Error sending notifications", err.Error())
	}

	return nil
}

package notifications

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

var notificationHandlerDomain string

func SetNotificationHandlerDomain(domain string) {
	notificationHandlerDomain = domain
}
func NotifyUser(ctx context.Context, userID uuid.UUID, key string, data map[string]string) error {
	url := fmt.Sprintf("%s/notify/user/%s", notificationHandlerDomain, userID)

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer([]byte(fmt.Sprintf(`{"key": "%s", "data": %v}`, key, data))))
	if err != nil {
		fmt.Println("Error creating request: ", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error publishing event: ", err)
		return err
	}

	return nil
}

func NotifyTopic(ctx context.Context, topic string, key string, data any) error {
	url := fmt.Sprintf("%s/notify/topic/%s", notificationHandlerDomain, topic)

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer([]byte(fmt.Sprintf(`{"key": "%s", "data": %v}`, key, data))))
	if err != nil {
		fmt.Println("Error creating request: ", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error publishing event: ", err)
		return err
	}

	return nil
}

func SubscribeToTopic(ctx context.Context, topic string, userID uuid.UUID) error {
	url := fmt.Sprintf("%s/subscribe/topic/%s/user/%s", notificationHandlerDomain, topic, userID)

	req, err := http.NewRequestWithContext(ctx, "POST", url, nil)
	if err != nil {
		fmt.Println("Error creating request: ", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error subscribing to topic: ", err)
		return err
	}

	return nil
}

func UnsubscribeFromTopic(ctx context.Context, topic string, userID uuid.UUID) error {
	url := fmt.Sprintf("%s/unsubscribe/topic/%s/user/%s", notificationHandlerDomain, topic, userID)

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		fmt.Println("Error creating request: ", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error unsubscribing from topic: ", err)
		return err
	}

	return nil
}

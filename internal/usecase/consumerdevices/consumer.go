package consumerdevices

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Melodia-IS2/melodia-events/internal/domain/entities"
	"github.com/Melodia-IS2/melodia-events/internal/domain/repositories"
	"github.com/Melodia-IS2/melodia-events/pkg/suscriber/kafka"
)

type ConsumerUserDevices struct {
	DeviceRepository repositories.DevicesRepository
}

func (c *ConsumerUserDevices) ConsumeBatch(ctx context.Context, topic string, msgs []kafka.BatchMessage) error {
	fmt.Println("-----------------------------------")
	fmt.Println("Consuming batch", topic, len(msgs))
	fmt.Println("-----------------------------------")

	insertMap := make(map[string]entities.Device)
	deleteMap := make(map[string]entities.Device)

	for _, msg := range msgs {
		var message Message
		err := json.Unmarshal(msg.Value, &message)
		if err != nil {
			log.Println("error unmarshalling message", err)
			continue
		}

		key := message.DeviceToken

		if key == "" {
			continue
		}

		switch msg.Key {

		case KeyLogin, KeyCreate:
			delete(deleteMap, key)

			insertMap[key] = entities.Device{
				UserID:      message.UserID,
				DeviceToken: message.DeviceToken,
			}

			fmt.Println("inserting device", message.UserID, message.DeviceToken)

		case KeyLogout:
			delete(insertMap, key)

			deleteMap[key] = entities.Device{
				UserID:      message.UserID,
				DeviceToken: message.DeviceToken,
			}

			fmt.Println("deleting device", message.UserID, message.DeviceToken)
		}
	}

	insertedDevices := make([]entities.Device, 0, len(insertMap))
	for _, d := range insertMap {
		insertedDevices = append(insertedDevices, d)
	}

	deletedDevices := make([]entities.Device, 0, len(deleteMap))
	for _, d := range deleteMap {
		deletedDevices = append(deletedDevices, d)
	}

	fmt.Println("inserting devices", len(insertedDevices))
	fmt.Println("deleting devices", len(deletedDevices))

	for _, d := range insertedDevices {
		err := c.DeviceRepository.Register(ctx, d)
		if err != nil {
			log.Println("error registering device", err)
			continue
		}
	}

	for _, d := range deletedDevices {
		err := c.DeviceRepository.Delete(ctx, d)
		if err != nil {
			log.Println("error deleting device", err)
			continue
		}
	}

	return nil
}

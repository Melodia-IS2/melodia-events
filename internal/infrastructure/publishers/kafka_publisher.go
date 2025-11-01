package publishers

import (
	"context"
	"encoding/json"

	"github.com/Melodia-IS2/melodia-events/internal/domain/entities"

	"github.com/segmentio/kafka-go"
)

type KafkaEventPublisher struct {
	Writer *kafka.Writer
}

func (p *KafkaEventPublisher) Publish(ctx context.Context, event *entities.Event) error {
	data, _ := json.Marshal(event.Payload)

	return p.Writer.WriteMessages(ctx, kafka.Message{
		Topic: event.Topic,
		Key:   []byte(event.Key),
		Value: data,
	})
}

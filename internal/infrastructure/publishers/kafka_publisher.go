package publishers

import (
	"context"
	"encoding/json"

	"melodia-events/internal/domain/entities"

	"github.com/segmentio/kafka-go"
)

type KafkaEventPublisher struct {
	Writer *kafka.Writer
}

func (p *KafkaEventPublisher) Publish(ctx context.Context, event *entities.Event) error {
	data, _ := json.Marshal(event)

	return p.Writer.WriteMessages(ctx, kafka.Message{
		Topic: event.Topic,
		Key:   []byte(event.Topic),
		Value: data,
	})
}

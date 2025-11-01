package kafka

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

type MessageHandler func(ctx context.Context, topic string, key string, msg []byte) error

type Config struct {
	Brokers []string
	GroupID string
	Topic   string
}

type SingleConsumer struct {
	reader  *kafka.Reader
	handler MessageHandler
}

func NewSingleConsumer(cfg Config, handler MessageHandler) *SingleConsumer {
	return &SingleConsumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:  cfg.Brokers,
			GroupID:  cfg.GroupID,
			Topic:    cfg.Topic,
			MinBytes: 10e3,
			MaxBytes: 10e6,
		}),
		handler: handler,
	}
}

func (c *SingleConsumer) Start(ctx context.Context) error {
	log.Printf("SingleConsumer listening to topic: %s", c.reader.Config().Topic)
	for {
		m, err := c.reader.ReadMessage(ctx)
		if err != nil {
			if ctx.Err() != nil {
				return nil
			}
			log.Printf("Error reading message: %v", err)
			time.Sleep(time.Second)
			continue
		}

		if err := c.handler(ctx, m.Topic, string(m.Key), m.Value); err != nil {
			log.Printf("Error processing message: %v", err)
		}
	}
}

func (c *SingleConsumer) Close() error {
	return c.reader.Close()
}

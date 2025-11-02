package kafka

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

type BatchMessage struct {
	Key   string
	Value []byte
}

type BatchMessageHandler func(ctx context.Context, topic string, msgs []BatchMessage) error

type BatchConfig struct {
	Brokers      []string
	GroupID      string
	Topic        string
	BatchSize    int
	BatchTimeout time.Duration
}

type BatchConsumer struct {
	reader  *kafka.Reader
	handler BatchMessageHandler
	cfg     BatchConfig
}

func NewBatchConsumer(cfg BatchConfig, handler BatchMessageHandler) *BatchConsumer {
	return &BatchConsumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:  cfg.Brokers,
			GroupID:  cfg.GroupID,
			Topic:    cfg.Topic,
			MinBytes: 10e3,
			MaxBytes: 10e6,
		}),
		handler: handler,
		cfg:     cfg,
	}
}

func (c *BatchConsumer) Start(ctx context.Context) error {
	log.Printf("Batch consumer listening to topic: %s", c.reader.Config().Topic)

	msgCh := make(chan kafka.Message)
	errCh := make(chan error)

	go func() {
		for {
			m, err := c.reader.ReadMessage(ctx)
			if err != nil {
				if ctx.Err() != nil {
					return
				}
				errCh <- err
				continue
			}
			msgCh <- m
		}
	}()

	var (
		batch []BatchMessage
	)
	timer := time.NewTimer(c.cfg.BatchTimeout)
	defer timer.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil

		case err := <-errCh:
			log.Printf("Error reading message: %v", err)

		case m := <-msgCh:
			batch = append(batch, BatchMessage{Key: string(m.Key), Value: m.Value})

			if len(batch) >= c.cfg.BatchSize {
				if err := c.flushBatch(ctx, c.reader.Config().Topic, &batch); err != nil {
					log.Printf("Error processing batch (full): %v", err)
				}
				batch = batch[:0]
				if !timer.Stop() {
					<-timer.C
				}
				timer.Reset(c.cfg.BatchTimeout)
			}

		case <-timer.C:
			if len(batch) > 0 {
				if err := c.flushBatch(ctx, c.reader.Config().Topic, &batch); err != nil {
					log.Printf("Error processing batch (timeout): %v", err)
				}
				batch = batch[:0]
			}
			timer.Reset(c.cfg.BatchTimeout)
		}
	}
}

func (c *BatchConsumer) flushBatch(ctx context.Context, topic string, batch *[]BatchMessage) error {
	err := c.handler(ctx, topic, *batch)
	*batch = (*batch)[:0]
	return err
}

func (c *BatchConsumer) Close() error {
	return c.reader.Close()
}

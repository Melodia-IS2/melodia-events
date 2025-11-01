package kafka

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

type BatchMessageHandler func(ctx context.Context, topic string, key string, msgs [][]byte) error

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

	var (
		batch   [][]byte
		lastKey string
		timer   = time.NewTimer(c.cfg.BatchTimeout)
	)
	defer timer.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil

		case <-timer.C:
			if len(batch) > 0 {
				if err := c.flushBatch(ctx, c.reader.Config().Topic, lastKey, &batch); err != nil {
					log.Printf("Error processing batch (timeout): %v", err)
				}
				batch = batch[:0]
			}
			timer.Reset(c.cfg.BatchTimeout)

		default:
			m, err := c.reader.ReadMessage(ctx)
			if err != nil {
				if ctx.Err() != nil {
					return nil
				}
				log.Printf("Error reading message: %v", err)
				time.Sleep(time.Second)
				continue
			}

			batch = append(batch, m.Value)
			lastKey = string(m.Key)

			if len(batch) >= c.cfg.BatchSize {
				if err := c.flushBatch(ctx, m.Topic, lastKey, &batch); err != nil {
					log.Printf("Error processing batch: %v", err)
				}
				batch = batch[:0]
				timer.Reset(c.cfg.BatchTimeout)
			}
		}
	}
}

func (c *BatchConsumer) flushBatch(ctx context.Context, topic, key string, batch *[][]byte) error {
	err := c.handler(ctx, topic, key, *batch)
	*batch = (*batch)[:0]
	return err
}

func (c *BatchConsumer) Close() error {
	return c.reader.Close()
}

package scheduler

import (
	"context"
	"log"
	"melodia-events/internal/domain/repositories"
	"time"
)

type Worker struct {
	EventRepository repositories.EventRepository
	EventPublisher  repositories.EventPublisher
	Interval        time.Duration
}

func (w *Worker) Start() error {
	if w.Interval <= 0 {
		w.Interval = 2 * time.Second
	}
	ctx := context.Background()
	ticker := time.NewTicker(w.Interval)
	defer ticker.Stop()

	for {
		if err := w.processDue(ctx); err != nil {
			log.Printf("worker process error: %v", err)
		}
		<-ticker.C
	}
}

func (w *Worker) Stop() error {
	return nil
}

func (w *Worker) processDue(ctx context.Context) error {
	now := time.Now()
	events, err := w.EventRepository.FindDueUnpublished(ctx, now)
	if err != nil {
		return err
	}
	for _, ev := range events {
		if err := w.EventPublisher.Publish(ctx, ev); err != nil {
			log.Printf("failed to publish event %v: %v", ev.ID, err)
			continue
		}
		if err := w.EventRepository.MarkPublished(ctx, ev.ID); err != nil {
			log.Printf("failed to mark event %v as published: %v", ev.ID, err)
			continue
		}
	}
	return nil
}

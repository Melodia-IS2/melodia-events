package persistence

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Melodia-IS2/melodia-events/internal/domain/entities"
	"github.com/redis/go-redis/v9"
)

type RedisEventRepository struct {
	Rdb *redis.Client
	Key string
}

func (r *RedisEventRepository) Schedule(ctx context.Context, e entities.Event) error {
	data, err := json.Marshal(e)
	if err != nil {
		return err
	}
	score := float64(e.PublishAfter.Unix())
	return r.Rdb.ZAdd(ctx, r.Key, redis.Z{Score: score, Member: data}).Err()
}

func (r *RedisEventRepository) FetchDueEvents(ctx context.Context, limit int64) ([]entities.Event, error) {
	now := float64(time.Now().Unix())

	results, err := r.Rdb.ZRangeByScore(ctx, r.Key, &redis.ZRangeBy{
		Min:    "-inf",
		Max:    fmt.Sprintf("%f", now),
		Offset: 0,
		Count:  limit,
	}).Result()
	if err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return nil, nil
	}

	_, err = r.Rdb.ZRem(ctx, r.Key, results).Result()
	if err != nil {
		return nil, err
	}

	events := make([]entities.Event, 0, len(results))
	for _, r := range results {
		var e entities.Event
		if err := json.Unmarshal([]byte(r), &e); err == nil {
			events = append(events, e)
		}
	}
	return events, nil
}

func (r *RedisEventRepository) FindAll(ctx context.Context) ([]entities.Event, error) {
	results, err := r.Rdb.ZRange(ctx, r.Key, 0, -1).Result()
	if err != nil {
		return nil, err
	}
	events := make([]entities.Event, 0, len(results))
	for _, r := range results {
		var e entities.Event
		if err := json.Unmarshal([]byte(r), &e); err == nil {
			events = append(events, e)
		}
	}
	return events, nil
}

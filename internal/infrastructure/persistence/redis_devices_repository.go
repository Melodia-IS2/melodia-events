package persistence

import (
	"context"
	"encoding/json"

	"github.com/Melodia-IS2/melodia-events/internal/domain/entities"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type RedisDevicesRepository struct {
	Rdb *redis.Client
	Key string
}

func (r *RedisDevicesRepository) keyForUser(userID uuid.UUID) string {
	return r.Key + ":" + userID.String()
}

func (r *RedisDevicesRepository) Register(ctx context.Context, device entities.Device) error {
	data, err := json.Marshal(device)
	if err != nil {
		return err
	}

	return r.Rdb.SAdd(ctx, r.keyForUser(device.UserID), data).Err()
}

func (r *RedisDevicesRepository) Delete(ctx context.Context, device entities.Device) error {
	data, err := json.Marshal(device)
	if err != nil {
		return err
	}

	return r.Rdb.SRem(ctx, r.keyForUser(device.UserID), data).Err()
}

func (r *RedisDevicesRepository) FetchByUserIDs(ctx context.Context, ids []uuid.UUID) ([]entities.Device, error) {
	devices := []entities.Device{}

	for _, id := range ids {
		result, err := r.Rdb.SMembers(ctx, r.keyForUser(id)).Result()
		if err != nil {
			return nil, err
		}

		for _, item := range result {
			var d entities.Device
			if err := json.Unmarshal([]byte(item), &d); err == nil {
				devices = append(devices, d)
			}
		}
	}

	return devices, nil
}

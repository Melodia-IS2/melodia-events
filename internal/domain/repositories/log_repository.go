package repositories

import (
	"context"

	"github.com/Melodia-IS2/melodia-events/internal/domain/entities"
	"github.com/Melodia-IS2/melodia-go-utils/pkg/logger"
)

type LogRepository interface {
	Log(ctx context.Context, log *logger.Log) error
	FindAll(ctx context.Context) ([]*logger.Log, error)
	Search(ctx context.Context, search entities.LogSearch) ([]*logger.Log, error)
}

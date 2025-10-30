package repositories

import (
	"context"

	"github.com/Melodia-IS2/melodia-go-utils/pkg/logger"
)

type LogRepository interface {
	Log(ctx context.Context, log *logger.Log) error
	FindAll(ctx context.Context) ([]*logger.Log, error)
}

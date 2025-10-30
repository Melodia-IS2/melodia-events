package getlogs

import (
	"context"

	"github.com/Melodia-IS2/melodia-events/internal/domain/repositories"
	"github.com/Melodia-IS2/melodia-go-utils/pkg/logger"
)

type GetLogs interface {
	Execute(ctx context.Context) ([]*logger.Log, error)
}

type GetLogsImpl struct {
	LogRepository repositories.LogRepository
}

func (u *GetLogsImpl) Execute(ctx context.Context) ([]*logger.Log, error) {
	return u.LogRepository.FindAll(ctx)
}

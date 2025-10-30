package createlog

import (
	"context"

	"github.com/Melodia-IS2/melodia-events/internal/domain/repositories"
	"github.com/Melodia-IS2/melodia-go-utils/pkg/logger"
)

type CreateLog interface {
	Execute(ctx context.Context, log logger.Log) error
}

type CreateLogImpl struct {
	LogRepository repositories.LogRepository
}

func (u *CreateLogImpl) Execute(ctx context.Context, log logger.Log) (err error) {
	if err := u.LogRepository.Log(ctx, &log); err != nil {
		print(err.Error())
		return err
	}

	return nil
}

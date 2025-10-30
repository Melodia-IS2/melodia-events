package createlog

import (
	"net/http"

	"github.com/Melodia-IS2/melodia-go-utils/pkg/errors"
	httpUtils "github.com/Melodia-IS2/melodia-go-utils/pkg/http"
	"github.com/Melodia-IS2/melodia-go-utils/pkg/logger"
	"github.com/Melodia-IS2/melodia-go-utils/pkg/router"
)

type CreateLogHandler struct {
	CreateLogUC CreateLog
}

func (handler *CreateLogHandler) Register(rt *router.Router) {
	rt.Post("/logs", handler.createLog)
}

func (handler *CreateLogHandler) createLog(w http.ResponseWriter, r *http.Request) error {
	log, err := httpUtils.ParseBody[logger.Log](r)
	if err != nil {
		return errors.NewBadRequestError(err.Error())
	}

	err = handler.CreateLogUC.Execute(r.Context(), log)

	if err != nil {
		return err
	}

	router.NoContent(w)

	return nil
}

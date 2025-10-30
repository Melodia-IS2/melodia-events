package getlogs

import (
	"net/http"

	"github.com/Melodia-IS2/melodia-go-utils/pkg/router"
)

type GetLogsHandler struct {
	GetLogsUC GetLogs
}

func (handler *GetLogsHandler) Register(rt *router.Router) {
	rt.Get("/logs", handler.getLogs)
}

func (handler *GetLogsHandler) getLogs(w http.ResponseWriter, r *http.Request) error {
	logs, err := handler.GetLogsUC.Execute(r.Context())
	if err != nil {
		return err
	}

	router.Ok(w, logs)

	return nil
}

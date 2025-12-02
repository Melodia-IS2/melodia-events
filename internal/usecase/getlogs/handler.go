package getlogs

import (
	"net/http"

	"github.com/Melodia-IS2/melodia-events/internal/domain/entities"
	"github.com/Melodia-IS2/melodia-go-utils/pkg/router"
)

type GetLogsHandler struct {
	GetLogsUC GetLogs
}

func (handler *GetLogsHandler) Register(rt *router.Router) {
	rt.Get("/logs", handler.getLogs)
	rt.Get("/logs/search", handler.getLogsSearch)
}

func (handler *GetLogsHandler) getLogs(w http.ResponseWriter, r *http.Request) error {
	logs, err := handler.GetLogsUC.Execute(r.Context())
	if err != nil {
		return err
	}

	router.Ok(w, logs)

	return nil
}

func (handler *GetLogsHandler) getLogsSearch(w http.ResponseWriter, r *http.Request) (err error) {
	search, err := new(entities.LogSearch).PopulateFromRequest(r)
	if err != nil {
		return err
	}

	logs, err := handler.GetLogsUC.Search(r.Context(), search)
	if err != nil {
		return err
	}

	router.Ok(w, logs)

	return nil
}

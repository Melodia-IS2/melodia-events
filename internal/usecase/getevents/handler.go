package getevents

import (
	"net/http"

	"github.com/Melodia-IS2/melodia-go-utils/pkg/router"
)

type GetEventsHandler struct {
	GetEventsUC GetEvents
}

func (handler *GetEventsHandler) Register(rt *router.Router) {
	rt.Get("/events", handler.getEvents)
}

func (handler *GetEventsHandler) getEvents(w http.ResponseWriter, r *http.Request) error {
	events, err := handler.GetEventsUC.Execute(r.Context())
	if err != nil {
		return err
	}

	router.Ok(w, events)

	return nil
}

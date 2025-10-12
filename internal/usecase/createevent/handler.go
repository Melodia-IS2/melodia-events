package createevent

import (
	"net/http"

	"github.com/Melodia-IS2/melodia-go-utils/pkg/errors"
	httpUtils "github.com/Melodia-IS2/melodia-go-utils/pkg/http"
	"github.com/Melodia-IS2/melodia-go-utils/pkg/router"
)

type CreateEventHandler struct {
	CreateEventUC CreateEvent
}

func (handler *CreateEventHandler) Register(rt *router.Router) {
	rt.Post("/event", handler.createEvent)
}

func (handler *CreateEventHandler) createEvent(w http.ResponseWriter, r *http.Request) error {
	req, err := httpUtils.ParseBody[CreateEventRequest](r)
	if err != nil {
		return errors.NewBadRequestError(err.Error())
	}

	event, err := req.ToDomain()
	if err != nil {
		return errors.NewBadRequestError(err.Error())
	}

	err = handler.CreateEventUC.Execute(r.Context(), event)

	if err != nil {
		return err
	}

	router.NoContent(w)

	return nil
}

package getnotifications

import (
	"net/http"

	"github.com/Melodia-IS2/melodia-go-utils/pkg/ctx"
	"github.com/Melodia-IS2/melodia-go-utils/pkg/errors"
	"github.com/Melodia-IS2/melodia-go-utils/pkg/router"
)

type GetNotificationsHandler struct {
	GetNotificationsUC GetNotifications
}

func (handler *GetNotificationsHandler) Register(rt *router.Router) {
	rt.Get("/notifications", handler.getNotifications)
}

func (handler *GetNotificationsHandler) getNotifications(w http.ResponseWriter, r *http.Request) error {
	n, err := router.GetQueryParam[uint](r, "n")
	if err != nil {
		return errors.NewBadRequestError("invalid n")
	}

	defaultN := uint(10)
	if n == nil {
		n = &defaultN
	}

	userID, err := ctx.GetUserID(r.Context())
	if err != nil {
		return errors.NewUnauthorizedError("unauthorized")
	}

	notifications, err := handler.GetNotificationsUC.Execute(r.Context(), *n, userID)
	if err != nil {
		return err
	}

	router.Ok(w, notifications)

	return nil
}

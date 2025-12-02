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

// GetNotifications godoc
// @Summary Get the notifications for a user
// @Tags notifications
// @Accept  json
// @Produce  json
// @Param n query int false "Number of notifications to return" default(10)
// @Success 200 {object} []entities.Notification "Notifications"
// @Failure 400 {object} errors.Error "Bad request error"
// @Failure 500 {object} errors.Error "Internal server error"
// @Router /notifications [get]
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

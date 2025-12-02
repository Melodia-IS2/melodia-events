package marknotificaitonasread

import (
	"net/http"

	"github.com/Melodia-IS2/melodia-go-utils/pkg/errors"
	"github.com/Melodia-IS2/melodia-go-utils/pkg/router"
	"github.com/google/uuid"
)

type MarkNotificationAsReadHandler struct {
	MarkNotificationAsReadUC MarkNotificationAsRead
}

func (handler *MarkNotificationAsReadHandler) Register(rt *router.Router) {
	rt.Put("/notifications/{notification_id}/read", handler.markNotificationAsRead)
}

// MarkNotificationAsRead godoc
// @Summary Mark a notification as read
// @Tags notifications
// @Accept  json
// @Produce  json
// @Param notification_id path string true "Notification ID"
// @Success 204 "Notification marked as read"
// @Failure 400 {object} errors.Error "Bad request error"
// @Failure 500 {object} errors.Error "Internal server error"
// @Router /notifications/{notification_id}/read [put]
func (handler *MarkNotificationAsReadHandler) markNotificationAsRead(w http.ResponseWriter, r *http.Request) error {
	notificationID, err := router.GetUrlParam[uuid.UUID](r, "notification_id")
	if err != nil {
		return errors.NewBadRequestError("invalid notification_id")
	}

	err = handler.MarkNotificationAsReadUC.Execute(r.Context(), notificationID)
	if err != nil {
		return err
	}

	router.NoContent(w)

	return nil
}

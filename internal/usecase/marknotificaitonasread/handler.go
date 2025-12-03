package marknotificaitonasread

import (
	"net/http"

	"github.com/Melodia-IS2/melodia-go-utils/pkg/auth"
	"github.com/Melodia-IS2/melodia-go-utils/pkg/ctx"
	"github.com/Melodia-IS2/melodia-go-utils/pkg/errors"
	"github.com/Melodia-IS2/melodia-go-utils/pkg/router"
	"github.com/google/uuid"
)

type MarkNotificationAsReadHandler struct {
	MarkNotificationAsReadUC MarkNotificationAsRead
	AuthMiddleware           auth.AuthMiddleware
}

func (handler *MarkNotificationAsReadHandler) Register(rt *router.Router) {
	rt.Put("/notifications/{notification_id}/read", handler.
		AuthMiddleware.
		NewBuilder().
		WithRol(ctx.ContextRolUser, ctx.ContextRolArtist).
		WithState(ctx.ContextStateActive).
		Build(handler.markNotificationAsRead))
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

	userID, err := ctx.GetUserID(r.Context())
	if err != nil {
		return errors.NewUnauthorizedError("unauthorized")
	}

	err = handler.MarkNotificationAsReadUC.Execute(r.Context(), notificationID, userID)
	if err != nil {
		return err
	}

	router.NoContent(w)

	return nil
}

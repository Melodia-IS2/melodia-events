package subnotifytopic

import (
	"net/http"

	"github.com/Melodia-IS2/melodia-go-utils/pkg/errors"
	"github.com/Melodia-IS2/melodia-go-utils/pkg/router"
	"github.com/google/uuid"
)

type SubNotifyTopicHandler struct {
	SubNotifyTopicUC SubNotifyTopic
}

func (handler *SubNotifyTopicHandler) Register(rt *router.Router) {
	rt.Post("/subscribe/topic/{topic}/user/{user_id}", handler.subNotifyTopic)
	rt.Delete("/unsubscribe/topic/{topic}/user/{user_id}", handler.unsubNotifyTopic)
}

func (handler *SubNotifyTopicHandler) subNotifyTopic(w http.ResponseWriter, r *http.Request) error {
	topic, err := router.GetUrlParam[string](r, "topic")
	if err != nil {
		return errors.NewBadRequestError("invalid topic")
	}

	userID, err := router.GetUrlParam[uuid.UUID](r, "user_id")
	if err != nil {
		return errors.NewBadRequestError("invalid user_id")
	}

	return handler.SubNotifyTopicUC.SubNotifyTopic(r.Context(), userID, topic)
}

func (handler *SubNotifyTopicHandler) unsubNotifyTopic(w http.ResponseWriter, r *http.Request) error {
	topic, err := router.GetUrlParam[string](r, "topic")
	if err != nil {
		return errors.NewBadRequestError("invalid topic")
	}

	userID, err := router.GetUrlParam[uuid.UUID](r, "user_id")
	if err != nil {
		return errors.NewBadRequestError("invalid user_id")
	}

	return handler.SubNotifyTopicUC.UnsubNotifyTopic(r.Context(), userID, topic)
}

package notify

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/Melodia-IS2/melodia-go-utils/pkg/errors"
	httpUtils "github.com/Melodia-IS2/melodia-go-utils/pkg/http"
	"github.com/Melodia-IS2/melodia-go-utils/pkg/router"
	"github.com/google/uuid"
)

type NotifyHandler struct {
	NotifyUC Notify
}

func (handler *NotifyHandler) Register(rt *router.Router) {
	rt.Post("/notify/user/{user_id}", handler.notifyUser)
	rt.Post("/notify/users", handler.notifyUsers)
	rt.Post("/notify/topic/{topic}", handler.notifyTopic)
}

func (handler *NotifyHandler) notifyUser(w http.ResponseWriter, r *http.Request) error {
	userID, err := router.GetUrlParam[uuid.UUID](r, "user_id")
	if err != nil {
		return errors.NewBadRequestError("invalid user_id")
	}

	req, err := httpUtils.ParseBody[NotifyRequest](r)
	if err != nil {
		return errors.NewBadRequestError("invalid request")
	}

	return handler.NotifyUC.NotifyUser(r.Context(), userID, req.Key, req.Data)
}

func (handler *NotifyHandler) notifyTopic(w http.ResponseWriter, r *http.Request) error {
	topic, err := router.GetUrlParam[string](r, "topic")
	if err != nil {
		return errors.NewBadRequestError("invalid topic")
	}

	req, err := httpUtils.ParseBody[NotifyRequest](r)
	if err != nil {
		return errors.NewBadRequestError("invalid request")
	}

	return handler.NotifyUC.NotifyTopic(r.Context(), topic, req.Key, req.Data)
}

func (handler *NotifyHandler) notifyUsers(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("NotifyUsers")
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading body:", err.Error())
	} else {
		fmt.Println("Raw body received:", string(bodyBytes))
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	}
	req, err := httpUtils.ParseBody[NotifyUsersRequest](r)
	if err != nil {
		fmt.Println("Error parsing request", err.Error())
		return errors.NewBadRequestError("invalid request")
	}

	fmt.Println("Request", req)

	return handler.NotifyUC.NotifyUsers(r.Context(), req.UserIDs, req.Key, req.Data)
}

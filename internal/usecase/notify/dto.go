package notify

import "github.com/google/uuid"

type NotifyRequest struct {
	Key  string            `json:"key"`
	Data map[string]string `json:"data"`
}

type NotifyUsersRequest struct {
	UserIDs []uuid.UUID       `json:"user_ids"`
	Key     string            `json:"key"`
	Data    map[string]string `json:"data"`
}

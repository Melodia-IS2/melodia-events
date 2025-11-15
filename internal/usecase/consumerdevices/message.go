package consumerdevices

import "github.com/google/uuid"

type Message struct {
	UserID      uuid.UUID `json:"user_id"`
	DeviceToken string    `json:"device_token"`
}

const (
	KeyCreate = "CREATE"
	KeyLogin  = "LOGIN"
	KeyLogout = "LOGOUT"
)

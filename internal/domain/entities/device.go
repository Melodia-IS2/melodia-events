package entities

import "github.com/google/uuid"

type Device struct {
	UserID      uuid.UUID
	DeviceToken string
}

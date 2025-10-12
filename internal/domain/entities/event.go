package entities

import (
	"time"

	"github.com/google/uuid"
)

type ErrorLevel int

const (
	NoErrorLevel      ErrorLevel = 0
	ErrorLevelInfo    ErrorLevel = 1
	ErrorLevelWarning ErrorLevel = 2
	ErrorLevelError   ErrorLevel = 3
)

type Event struct {
	ID          uuid.UUID
	Topic       string
	Title       string
	Description string
	Service     string
	Payload     string
	ErrorLevel  ErrorLevel
	Token       string
	CreatedAt   time.Time
}

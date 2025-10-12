package events

import (
	"time"
)

type ErrorLevel int

const (
	NoErrorLevel      ErrorLevel = 0
	ErrorLevelInfo    ErrorLevel = 1
	ErrorLevelWarning ErrorLevel = 2
	ErrorLevelError   ErrorLevel = 3
)

type Event struct {
	Topic       string     `json:"topic"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Service     string     `json:"service"`
	Payload     string     `json:"payload"`
	ErrorLevel  ErrorLevel `json:"error_level"`
	Token       string     `json:"token"`
	CreatedAt   time.Time  `json:"created_at"`
}

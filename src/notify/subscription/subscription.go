package subscription

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Subscription is a structure that matches the json content for REST calls
type Subscription struct {
	ID        uuid.UUID `json:"id"`
	EventType string    `json:"event_type"`
	Context   string    `json:"context"`
	AccountID uuid.UUID `json:"account_id"`
	CreatedAt time.Time `json:"timestamp"`
}

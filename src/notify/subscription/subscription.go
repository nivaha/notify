package subscription

import "time"

// Subscription is a structure that matches the json content for REST calls
type Subscription struct {
	ID        string    `json:"id"`
	EventType string    `json:"event_type"`
	Context   string    `json:"context"`
	AccountID string    `json:"account_id"`
	CreatedAt time.Time `json:"timestamp"`
}

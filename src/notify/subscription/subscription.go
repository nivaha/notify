package subscription

import "time"

// Subscription is a structure that matches the json content for REST calls
type Subscription struct {
	ID                string    `json:"id"`
	EventType         string    `json:"event_type"`
	Context           string    `json:"context"`
	OriginalAccountID string    `json:"original_account_id"`
	CreatedAt         time.Time `json:"timestamp"`
}

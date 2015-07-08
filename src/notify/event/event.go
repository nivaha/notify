package event

import "time"

type Event struct {
	ID                string    `json:"id"`
	EventType         string    `json:"event_type"`
	Context           string    `json:"context"`
	OriginalAccountID string    `json:"original_account_id"`
	CreatedAt         time.Time `json:"timestamp"`
	Data              string    `json:"payload"`
}

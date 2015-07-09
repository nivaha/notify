package event

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Event struct {
	ID                uuid.UUID `json:"id"`
	EventType         string    `json:"event_type"`
	Context           string    `json:"context"`
	OriginalAccountID uuid.UUID `json:"original_account_id"`
	Data              string    `json:"payload"`
	CreatedAt         time.Time `json:"timestamp"`
}

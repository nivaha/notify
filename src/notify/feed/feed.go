package feed

import (
	"time"

	"notify/uuids"

	uuid "github.com/satori/go.uuid"
)

// Feed defines the data for an event feed
type Feed struct {
	ID        uuid.UUID   `json:"id"`
	Unread    uuids.UUIDs `json:"unread_events"`
	Events    uuids.UUIDs `json:"events"`
	CreatedAt time.Time   `json:"timestamp"`
}

package feed

import uuid "github.com/satori/go.uuid"

type EventIDs []uuid.UUID

type Feed struct {
	ID     uuid.UUID `json:"id"`
	Unread EventIDs  `json:"unread_events"`
	Events EventIDs  `json:"events"`
}

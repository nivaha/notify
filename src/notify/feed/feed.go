package feed

type EventIDs []string

type Feed struct {
	ID     string   `json:"id"`
	Unread EventIDs `json:"unread_events"`
	Events EventIDs `json:"events"`
}

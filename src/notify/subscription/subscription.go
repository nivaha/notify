package subscription

type AccountIDs []string

type Subscription struct {
	ID        string     `json:"id"`
	EventType string     `json:"event_type"`
	Context   string     `json:"context"`
	Accounts  AccountIDs `json:"accounts"`
}

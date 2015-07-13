package subscription

import (
	"encoding/json"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
)

func TestSubscription(t *testing.T) {
	var s Subscription
	var sJSON = []byte(`
    {
        "id": "5b2e1b6f-b3f7-4a94-8b05-aa441606d886",
        "event_type": "test_event_type",
        "context": "test_context",
        "account_id": "e4f05126-a0e8-437d-94f9-b553c959cdfb",
        "timestamp": "2015-07-10T10:50:34.437512Z"
    }`)

	if err := json.Unmarshal(sJSON, &s); err != nil {
		t.Error(err.Error())
	}

	sAccountID, _ := uuid.FromString("e4f05126-a0e8-437d-94f9-b553c959cdfb")
	sID, _ := uuid.FromString("5b2e1b6f-b3f7-4a94-8b05-aa441606d886")
	sTime, _ := time.Parse(time.RFC3339, "2015-07-10T10:50:34.437512Z")

	cases := []struct {
		label, actual, expected interface{}
	}{
		{"id", s.ID, sID},
		{"event type", s.EventType, "test_event_type"},
		{"context", s.Context, "test_context"},
		{"account id", s.AccountID, sAccountID},
		{"created at date", s.CreatedAt, sTime},
	}

	for _, c := range cases {
		if c.actual != c.expected {
			t.Errorf("Subscriber %v is %q, but expected %q", c.label, c.actual, c.expected)
		}
	}
}

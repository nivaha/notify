package event

import (
	"encoding/json"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
)

func TestEvent(t *testing.T) {
	var e Event
	var eJSON = []byte(`
    {
        "id": "5b2e1b6f-b3f7-4a94-8b05-aa441606d886",
        "event_type": "test_event_type",
        "context": "test_context",
        "original_account_id": "e4f05126-a0e8-437d-94f9-b553c959cdfb",
				"payload": "Инстинкт плюс возможность, равно прибыль.",
        "timestamp": "2015-07-10T10:50:34.437512Z"
    }`)

	if err := json.Unmarshal(eJSON, &e); err != nil {
		t.Error(err.Error())
	}

	eAccountID, _ := uuid.FromString("e4f05126-a0e8-437d-94f9-b553c959cdfb")
	eID, _ := uuid.FromString("5b2e1b6f-b3f7-4a94-8b05-aa441606d886")
	eTime, _ := time.Parse(time.RFC3339, "2015-07-10T10:50:34.437512Z")

	cases := []struct {
		label, actual, expected interface{}
	}{
		{"id", e.ID, eID},
		{"event type", e.EventType, "test_event_type"},
		{"context", e.Context, "test_context"},
		{"original account id", e.OriginalAccountID, eAccountID},
		{"created at date", e.CreatedAt, eTime},
		{"payload data", e.Data, "Инстинкт плюс возможность, равно прибыль."},
	}

	for _, c := range cases {
		if c.actual != c.expected {
			t.Errorf("Event %v is %q, but expected %q", c.label, c.actual, c.expected)
		}
	}
}

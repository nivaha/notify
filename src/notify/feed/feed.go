package feed

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

// EventUUIDs is a list of UUIDs
type EventUUIDs []uuid.UUID

// Feed defines the data for an event feed
type Feed struct {
	ID        uuid.UUID  `json:"id"`
	Unread    EventUUIDs `json:"unread_events"`
	Events    EventUUIDs `json:"events"`
	CreatedAt time.Time  `json:"timestamp"`
}

// Value converts EventUUIDs to a sql value
func (e EventUUIDs) Value() (driver.Value, error) {
	var uuids []string
	for i := range e {
		uuids = append(uuids, e[i].String())
	}

	result := fmt.Sprintf("{%v}", strings.Join(uuids, ","))

	return result, nil
}

// Scan converts an sql value to EventUUIDs
func (e *EventUUIDs) Scan(src interface{}) error {
	if src == nil {
		return nil
	}

	asBytes, ok := src.([]byte)
	if !ok {
		return error(errors.New("Scan source was not []bytes"))
	}

	var err error
	(*e), err = strToEventIDSlice(string(asBytes))

	return err
}

func strToEventIDSlice(str string) (EventUUIDs, error) {
	result := make(EventUUIDs, 0, 10)
	trimmed := strings.Trim(str, "{}")

	for _, t := range strings.Split(trimmed, ",") {
		id, err := uuid.FromString(t)
		if err != nil {
			return nil, err
		}

		result = append(result, id)
	}

	return result, nil
}

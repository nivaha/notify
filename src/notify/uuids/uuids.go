package uuids

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"

	uuid "github.com/satori/go.uuid"
)

// UUIDs is a list of UUIDs
type UUIDs []uuid.UUID

// Value converts UUIDs to a sql value
func (e UUIDs) Value() (driver.Value, error) {
	var uuids []string
	for i := range e {
		uuids = append(uuids, e[i].String())
	}

	result := fmt.Sprintf("{%v}", strings.Join(uuids, ","))

	return result, nil
}

// Scan converts an sql value to UUIDs
func (e *UUIDs) Scan(src interface{}) error {
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

func strToEventIDSlice(str string) (UUIDs, error) {
	result := make(UUIDs, 0, 10)
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

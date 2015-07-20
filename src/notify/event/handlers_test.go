package event

import (
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/julienschmidt/httprouter"
)

var testEvent = Event{
	EventType:         "test_type",
	Context:           "test_context",
	ID:                toID("e4f05126-a0e8-437d-94f9-b553c959cdfb"),
	OriginalAccountID: toID("5b2e1b6f-b3f7-4a94-8b05-aa441606d886"),
	CreatedAt:         toTime("2015-07-10T10:50:34.437512Z"),
	Data:              "Инстинкт плюс возможность, равно прибыль.",
}

// Mock out the Database dependency
type MockDatabase struct {
	t     *testing.T
	getID string
	e     Event
	elist []Event
	err   error
}

func (p MockDatabase) lookup(id string) (Event, error) {
	if id != p.getID {
		p.t.Errorf("Got a retrieval id of '%v', but expected '%v'", id, p.getID)
	}
	return p.e, p.err
}
func (p MockDatabase) list() ([]Event, error) {
	return p.elist, p.err
}

/* **** Test Cases **** */

// Test that the Index handler properly returns the list of retrieved events and correct http response code
func TestIndexHandler(t *testing.T) {
	elist := []Event{testEvent}

	h := Handler{
		db: MockDatabase{elist: elist},
	}
	req, w := newReqParams("GET")

	h.Index(w, req, httprouter.Params{})

	cases := []struct {
		label, actual, expected interface{}
	}{
		{"Response code", w.Code, 200},
		{"Response body contains context", strings.Contains(w.Body.String(), testEvent.Context), true},
		{"Response body contains event type", strings.Contains(w.Body.String(), testEvent.EventType), true},
		{"Response body contains data", strings.Contains(w.Body.String(), testEvent.Data), true},
		{"Response body contains id", strings.Contains(w.Body.String(), testEvent.ID.String()), true},
		{"Response body contains account id", strings.Contains(w.Body.String(), testEvent.OriginalAccountID.String()), true},
	}

	testCases(t, cases)
}

// Test that the Index handler properly returns the error message and correct http response code
func TestIndexHandlerWithErr(t *testing.T) {
	errMsg := "Bad things happened"
	h := Handler{
		db: MockDatabase{err: errors.New(errMsg)},
	}
	req, w := newReqParams("GET")

	h.Index(w, req, httprouter.Params{})

	cases := []struct {
		label, actual, expected interface{}
	}{
		{"Response code", w.Code, 500},
		{"Response body contains error message", strings.Contains(w.Body.String(), errMsg), true},
	}

	testCases(t, cases)
}

// Test that the Show handler properly returns the requested event by id and correct http response code
func TestShowHandler(t *testing.T) {
	id := "c79c54de-39ae-46b0-90e5-9f84c77f6974"
	params := httprouter.Params{
		httprouter.Param{Key: "id", Value: id},
	}

	h := Handler{
		db: MockDatabase{
			e:     testEvent,
			getID: id,
		},
	}

	req, w := newReqParams("GET")

	h.Show(w, req, params)

	cases := []struct {
		label, actual, expected interface{}
	}{
		{"Response code", w.Code, 200},
		{"Response body contains context", strings.Contains(w.Body.String(), testEvent.Context), true},
		{"Response body contains event type", strings.Contains(w.Body.String(), testEvent.EventType), true},
		{"Response body contains data", strings.Contains(w.Body.String(), testEvent.Data), true},
		{"Response body contains id", strings.Contains(w.Body.String(), testEvent.ID.String()), true},
		{"Response body contains account id", strings.Contains(w.Body.String(), testEvent.OriginalAccountID.String()), true},
	}

	testCases(t, cases)
}

// Test that the Index handler properly returns the error message and correct http response code
func TestShowHandlerWithErr(t *testing.T) {
	errMsg := "Bad things happened"
	id := "c79c54de-39ae-46b0-90e5-9f84c77f6974"
	params := httprouter.Params{
		httprouter.Param{Key: "id", Value: id},
	}
	h := Handler{
		db: MockDatabase{
			err:   errors.New(errMsg),
			getID: id,
		},
	}

	req, w := newReqParams("GET")

	h.Index(w, req, params)

	cases := []struct {
		label, actual, expected interface{}
	}{
		{"Response code", w.Code, 500},
		{"Response body contains error message", strings.Contains(w.Body.String(), errMsg), true},
	}

	testCases(t, cases)
}

/* **** Private **** */

func newReqParams(reqType string) (*http.Request, *httptest.ResponseRecorder) {
	req, _ := http.NewRequest(reqType, "/fake", strings.NewReader(""))
	w := httptest.NewRecorder()
	return req, w
}

func toID(param string) uuid.UUID {
	val, err := uuid.FromString(param)
	if err != nil {
		log.Fatal(err)
	}
	return val
}

func toTime(param string) time.Time {
	val, err := time.Parse(time.RFC3339, param)
	if err != nil {
		log.Fatal(err)
	}
	return val
}

func testCases(t *testing.T, cases []struct{ label, actual, expected interface{} }) {
	for _, c := range cases {
		if c.actual != c.expected {
			t.Errorf("%v is '%v', but expected '%v'", c.label, c.actual, c.expected)
		}
	}
}

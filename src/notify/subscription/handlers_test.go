package subscription

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/julienschmidt/httprouter"
)

// Mock out the Database dependency
type MockPersistance struct {
	t     *testing.T
	getID string
	s     Subscription
	slist []Subscription
	err   error
}

func (p MockPersistance) get(id string) (Subscription, error) {
	if id != p.getID {
		p.t.Errorf("Got a retrieval id of '%v', but expected '%v'", id, p.getID)
	}
	return p.s, p.err
}
func (p MockPersistance) list() ([]Subscription, error) {
	return p.slist, p.err
}
func (p MockPersistance) destroy(id string) (Subscription, error) {
	return p.s, p.err
}

// Test that the Index handler properly returns the list of retrieved subscriptions and correct http response code
func TestIndexHandler(t *testing.T) {
	slist := []Subscription{
		Subscription{
			EventType: "test_type",
			Context:   "test_context",
		},
	}

	h := Handler{
		db: MockPersistance{slist: slist},
	}
	req, w := newReqParams("GET")

	h.Index(w, req, httprouter.Params{})

	cases := []struct {
		label, actual, expected interface{}
	}{
		{"Response code", w.Code, 200},
		{"Response body contains context", strings.Contains(w.Body.String(), slist[0].Context), true},
		{"Response body contains event type", strings.Contains(w.Body.String(), slist[0].EventType), true},
	}

	testCases(t, cases)
}

// Test that the Index handler properly returns the error message and correct http response code
func TestIndexHandlerWithErr(t *testing.T) {
	errMsg := "Bad things happened"
	h := Handler{
		db: MockPersistance{err: errors.New(errMsg)},
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

// Test that the Show handler properly returns the requested subscription by id and correct http response code
func TestShowHandler(t *testing.T) {
	id := "c79c54de-39ae-46b0-90e5-9f84c77f6974"
	params := httprouter.Params{
		httprouter.Param{"id", id},
	}
	s := Subscription{
		EventType: "test_type",
		Context:   "test_context",
	}

	h := Handler{
		db: MockPersistance{
			s:     s,
			getID: id,
		},
	}

	req, w := newReqParams("GET")

	h.Show(w, req, params)

	cases := []struct {
		label, actual, expected interface{}
	}{
		{"Response code", w.Code, 200},
		{"Response body contains context", strings.Contains(w.Body.String(), s.Context), true},
		{"Response body contains event type", strings.Contains(w.Body.String(), s.EventType), true},
	}

	testCases(t, cases)
}

// Test that the Index handler properly returns the error message and correct http response code
func TestShowHandlerWithErr(t *testing.T) {
	errMsg := "Bad things happened"
	id := "c79c54de-39ae-46b0-90e5-9f84c77f6974"
	params := httprouter.Params{
		httprouter.Param{"id", id},
	}
	h := Handler{
		db: MockPersistance{
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

func testCases(t *testing.T, cases []struct{ label, actual, expected interface{} }) {
	for _, c := range cases {
		if c.actual != c.expected {
			t.Errorf("%v is '%v', but expected '%v'", c.label, c.actual, c.expected)
		}
	}
}

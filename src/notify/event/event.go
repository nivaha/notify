package event

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"notify/httpUtils"

	"github.com/julienschmidt/httprouter"
)

type Event struct {
	ID                string    `json:"id"`
	EventType         string    `json:"event_type"`
	Context           string    `json:"context"`
	OriginalAccountID string    `json:"original_account_id"`
	CreatedAt         time.Time `json:"timestamp"`
	Data              string    `json:"payload"`
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	events, err := list()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	fmt.Fprint(w, "index of events\n")
	for i := range events {
		fmt.Fprintf(w, "Event %v\n", events[i])
	}
}

func Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	e := Event{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&e)
	if err != nil {
		httpUtils.ErrorJSON(w, http.StatusBadRequest, "Could not decode request")
		return
	}

	err = e.insert()
	if err != nil {
		httpUtils.ErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	eJSON, _ := json.Marshal(e)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%v", eJSON)
}

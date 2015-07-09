package event

import (
	"encoding/json"
	"net/http"

	"notify/jsonUtils"

	"github.com/julienschmidt/httprouter"
)

// Index returns a JSON array of all events
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	events, err := list()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonUtils.Output(w, 200, events)
}

// Create constructs a new event from the data in the POST body
func Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	e := Event{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&e)
	if err != nil {
		jsonUtils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	err = e.insert()
	if err != nil {
		jsonUtils.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonUtils.Output(w, 201, e)
}

// Show returns the data for a speciic event as JSON
func Show(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	e, err := lookup(ps.ByName("id"))
	if err != nil {
		jsonUtils.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonUtils.Output(w, 200, e)
}

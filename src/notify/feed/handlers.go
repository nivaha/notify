package feed

import (
	"encoding/json"
	"net/http"

	"notify/jsonUtils"

	"github.com/julienschmidt/httprouter"
)

// Index returns a JSON array of all feeds
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	feeds, err := list()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonUtils.Output(w, 200, feeds)
}

// Create constructs a new feed from the data in the POST body
func Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	f := Feed{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&f)
	if err != nil {
		jsonUtils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	err = f.insert()
	if err != nil {
		jsonUtils.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonUtils.Output(w, 201, f)
}

// Show returns the data for a specific feed as JSON
func Show(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	f, err := lookup(ps.ByName("id"))
	if err != nil {
		jsonUtils.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonUtils.Output(w, 200, f)
}

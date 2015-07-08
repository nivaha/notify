package event

import (
	"encoding/json"
	"fmt"
	"net/http"

	"notify/httpUtils"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	events, err := list()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	w.Write([]byte("index of events\n"))
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
	w.Write(eJSON)
}

func Show(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	e, err := lookup(ps.ByName("id"))
	if err != nil {
		httpUtils.ErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	eJSON, _ := json.Marshal(e)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write(eJSON)
}

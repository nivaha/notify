package subscription

import (
	"encoding/json"
	"net/http"
	"notify/jsonUtils"

	"github.com/julienschmidt/httprouter"
)

// Create is a REST API for creating a new subscription, based on the JSON payload
func Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	subscription := Subscription{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&subscription)
	if err != nil {
		jsonUtils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	err = subscription.insert()
	if err != nil {
		jsonUtils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	jsonUtils.Output(w, 201, subscription)
}

// Index is a REST API for listing all registered subscriptions
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	subscriptions, err := list()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonUtils.Output(w, 200, subscriptions)
}

// Show is a REST API for listing a single subscription, found by id
func Show(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	subscription, err := get(p.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonUtils.Output(w, 200, subscription)
}

// Destroy is a REST API for destroying an subscription, based on the subscription id
func Destroy(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	subscription, err := destroy(p.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonUtils.Output(w, 200, subscription)
}

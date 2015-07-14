package subscription

import (
	"net/http"
	"notify/jsonUtils"

	"github.com/julienschmidt/httprouter"
)

type Handler struct {
	db Persist
}

func NewHandler() *Handler {
	h := Handler{
		db: Persistance{},
	}

	return &h
}

// Create is a REST API for creating a new subscription, based on the JSON payload
func (h *Handler) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	subscription := Subscription{}

	if err := jsonUtils.Decode(r.Body, &subscription); err != nil {
		jsonUtils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := subscription.insert(); err != nil {
		jsonUtils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	jsonUtils.Output(w, 201, subscription)
}

// Index is a REST API for listing all registered subscriptions
func (h *Handler) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	subscriptions, err := h.db.list()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonUtils.Output(w, 200, subscriptions)
}

// Show is a REST API for listing a single subscription, found by id
func (h *Handler) Show(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	subscription, err := h.db.get(p.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonUtils.Output(w, 200, subscription)
}

// Destroy is a REST API for destroying an subscription, based on the subscription id
func (h *Handler) Destroy(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	subscription, err := h.db.destroy(p.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonUtils.Output(w, 200, subscription)
}

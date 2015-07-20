package event

import (
	"net/http"

	"notify/jsonUtils"

	"github.com/julienschmidt/httprouter"
)

type Handler struct {
	db Database
}

func NewHandler() *Handler {
	h := Handler{
		db: PgDatabase{},
	}

	return &h
}

// Index returns a JSON array of all events
func (h *Handler) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	events, err := h.db.list()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonUtils.Output(w, 200, events)
}

// Create constructs a new event from the data in the POST body
func (h *Handler) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	e := Event{}

	if err := jsonUtils.Decode(r.Body, &e); err != nil {
		jsonUtils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := e.insert(); err != nil {
		jsonUtils.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonUtils.Output(w, 201, e)
}

// Show returns the data for a specific event as JSON
func (h *Handler) Show(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	e, err := h.db.lookup(ps.ByName("id"))
	if err != nil {
		jsonUtils.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonUtils.Output(w, 200, e)
}

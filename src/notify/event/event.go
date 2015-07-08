package event

import (
	"database/sql"
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

var myDB *sql.DB
var prepStmts struct {
	insert *sql.Stmt
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rows, err := myDB.Query("SELECT * FROM events")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	events := []Event{}

	for rows.Next() {
		var e Event
		var data string

		err := rows.Scan(&e.ID, &e.EventType, &e.Context, &e.OriginalAccountID, &e.CreatedAt, &data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		e.Data = data

		events = append(events, e)
	}

	err = rows.Err()
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

	_, err = prepStmts.insert.Exec(e.EventType, e.Context, e.OriginalAccountID)
	if err != nil {
		httpUtils.ErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	eJSON, _ := json.Marshal(e)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%v", eJSON)
}

func New(db *sql.DB) error {
	myDB = db

	err := prepareStatements()

	create := `CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
            CREATE TABLE IF NOT EXISTS events (
              id UUID PRIMARY KEY,
              event_type VARCHAR(64),
              context VARCHAR(64),
              original_account_id VARCHAR(64),
              created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
              data JSON DEFAULT '{}'::json)`

	_, err = db.Exec(create)

	return err
}

func prepareStatements() error {
	var err error

	prepStmts.insert, err = myDB.Prepare(`INSERT INTO events
    ( id,
      event_type,
      context,
      original_account_id
      )
    VALUES ( uuid_generate_v4(), $1, $2, $3 )
  `)

	return err
}

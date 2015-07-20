package event

import "database/sql"

var myDB *sql.DB
var prepStmts struct {
	list   *sql.Stmt
	lookup *sql.Stmt
	insert *sql.Stmt
}

// PgDatabase is a structure that implements the Database interface and interacts with Postgres db
type PgDatabase struct{}

// Database is an interface for retrieving Events
type Database interface {
	lookup(id string) (Event, error)
	list() ([]Event, error)
}

// CreateDB creates the events table if it does not yet exist
func CreateDB(db *sql.DB) error {
	myDB = db

	_, err := myDB.Exec(`
              CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
              CREATE TABLE IF NOT EXISTS events (
								id UUID PRIMARY KEY NOT NULL,
								event_type VARCHAR(64),
								context VARCHAR(64),
								original_account_id UUID,
								data JSON DEFAULT '{}'::json,
								created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
                )
              `)
	if err != nil {
		return err
	}

	err = prepareStatements()

	return err
}

func (p PgDatabase) list() ([]Event, error) {
	rows, err := prepStmts.list.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := []Event{}

	for rows.Next() {
		var e Event
		var data string

		err := rows.Scan(&e.ID, &e.EventType, &e.Context, &e.OriginalAccountID, &data, &e.CreatedAt)
		if err != nil {
			return nil, err
		}
		e.Data = data

		events = append(events, e)
	}

	return events, rows.Err()
}

func (p PgDatabase) lookup(id string) (Event, error) {
	var e Event
	var data string

	err := prepStmts.lookup.QueryRow(id).Scan(&e.ID, &e.EventType, &e.Context, &e.OriginalAccountID, &data, &e.CreatedAt)
	if err != nil {
		return Event{}, err
	}
	e.Data = data

	return e, nil
}

func (e *Event) insert() error {
	return prepStmts.insert.QueryRow(e.EventType, e.Context, e.OriginalAccountID.String()).Scan(&e.ID)
}

func prepareStatements() error {
	var err error

	prepStmts.list, err = myDB.Prepare(`SELECT * FROM events`)
	if err != nil {
		return err
	}

	prepStmts.lookup, err = myDB.Prepare(`SELECT * FROM events WHERE id = $1`)
	if err != nil {
		return err
	}

	prepStmts.insert, err = myDB.Prepare(`INSERT INTO events
		( id,
			event_type,
			context,
			original_account_id
		)
		VALUES ( uuid_generate_v4(), $1, $2, $3 )
		RETURNING id
  `)

	return err
}

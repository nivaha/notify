package event

import "database/sql"

var myDB *sql.DB
var prepStmts struct {
	insert *sql.Stmt
}

func CreateDB(db *sql.DB) error {
	myDB = db

	_, err := myDB.Exec(`
              CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
              CREATE TABLE IF NOT EXISTS events (
                id UUID PRIMARY KEY,
                event_type VARCHAR(64),
                context VARCHAR(64),
                original_account_id VARCHAR(64),
                created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                data JSON DEFAULT '{}'::json)
              `)
	if err != nil {
		return err
	}

	err = prepareStatements()

	return err
}

func list() ([]Event, error) {
	rows, err := myDB.Query("SELECT * FROM events")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := []Event{}

	for rows.Next() {
		var e Event
		var data string

		err := rows.Scan(&e.ID, &e.EventType, &e.Context, &e.OriginalAccountID, &e.CreatedAt, &data)
		if err != nil {
			return nil, err
		}
		e.Data = data

		events = append(events, e)
	}

	err = rows.Err()
	return events, err
}

func (e Event) insert() error {
	_, err := prepStmts.insert.Exec(e.EventType, e.Context, e.OriginalAccountID)

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

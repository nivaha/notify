package feed

import "database/sql"

var myDB *sql.DB
var prepStmts struct {
	list   *sql.Stmt
	lookup *sql.Stmt
	insert *sql.Stmt
}

// CreateDB creates the events table if it does not yet exist
func CreateDB(db *sql.DB) error {
	myDB = db

	_, err := myDB.Exec(`
              CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
              CREATE TABLE IF NOT EXISTS feeds (
								id UUID PRIMARY KEY NOT NULL,
                unread_events UUID ARRAY,
                events UUID ARRAY,
								created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
                )
              `)
	if err != nil {
		return err
	}

	err = prepareStatements()

	return err
}

func list() ([]Feed, error) {
	rows, err := prepStmts.list.Query()

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	feeds := []Feed{}

	for rows.Next() {
		var f Feed

		err := rows.Scan(&f.ID, &f.Unread, &f.Events, &f.CreatedAt)
		if err != nil {
			return nil, err
		}

		feeds = append(feeds, f)
	}

	err = rows.Err()
	return feeds, err
}

func lookup(id string) (Feed, error) {
	var f Feed

	err := prepStmts.lookup.QueryRow(id).Scan(&f.ID, &f.Unread, &f.Events, &f.CreatedAt)
	if err != nil {
		return Feed{}, err
	}

	return f, err
}

func (f *Feed) insert() error {
	err := prepStmts.insert.QueryRow(f.Unread, f.Events).Scan(&f.ID)

	return err
}

func prepareStatements() error {
	var err error

	prepStmts.list, err = myDB.Prepare(`SELECT * FROM feeds`)
	if err != nil {
		return err
	}

	prepStmts.lookup, err = myDB.Prepare(`SELECT * FROM feeds WHERE id = $1`)
	if err != nil {
		return err
	}

	prepStmts.insert, err = myDB.Prepare(`INSERT INTO feeds
		( id,
			unread_events,
			events
		)
		VALUES ( uuid_generate_v4(), $1, $2 )
		RETURNING id
  `)

	return err
}

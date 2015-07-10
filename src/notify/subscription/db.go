package subscription

import "database/sql"

var myDB *sql.DB
var prepStmts struct {
	insert  *sql.Stmt
	destroy *sql.Stmt
	get     *sql.Stmt
	list    *sql.Stmt
}

// CreateDB will setup the subscriptions table if it does not yet exist
func CreateDB(db *sql.DB) error {
	myDB = db

	if _, err := myDB.Exec(`
      CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
      CREATE TABLE IF NOT EXISTS subscriptions (
        id UUID PRIMARY KEY,
        event_type VARCHAR(64),
        context VARCHAR(64),
        account_id UUID,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP )
    `); err != nil {
		return err
	}

	err := prepareStatements()

	return err
}

func get(id string) (Subscription, error) {
	var subscription Subscription

	rows, err := prepStmts.get.Query(id)
	if err != nil {
		return subscription, err
	}

	defer rows.Close()

	if rows.Next() {
		if err := scan(rows, &subscription); err != nil {
			return subscription, err
		}
	}

	return subscription, nil
}

func list() ([]Subscription, error) {
	rows, err := prepStmts.list.Query()

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	subscriptions := []Subscription{}

	for rows.Next() {
		var subscription Subscription

		if err := scan(rows, &subscription); err != nil {
			return nil, err
		}

		subscriptions = append(subscriptions, subscription)
	}

	err = rows.Err()
	return subscriptions, err
}

func (subscription *Subscription) insert() error {
	err := prepStmts.insert.QueryRow(subscription.EventType, subscription.Context, subscription.AccountID.String()).Scan(&subscription.ID, &subscription.CreatedAt)

	return err
}

func destroy(id string) (Subscription, error) {
	subscription, err := get(id)
	if err != nil {
		return subscription, err
	}

	_, err = prepStmts.destroy.Query(id)
	return subscription, err
}

func scan(rows *sql.Rows, subscription *Subscription) error {
	return rows.Scan(&subscription.ID, &subscription.EventType, &subscription.Context, &subscription.AccountID, &subscription.CreatedAt)
}

func prepareStatements() error {
	var err error

	if prepStmts.insert, err = myDB.Prepare(`
			INSERT INTO subscriptions
	    ( id,
	      event_type,
	      context,
	      account_id
	    )
	    VALUES ( uuid_generate_v4(), $1, $2, $3 )
			RETURNING id, created_at
	  `); err != nil {
		return err
	}
	if prepStmts.destroy, err = myDB.Prepare(`
			DELETE
			FROM subscriptions
	    WHERE id = $1
	  `); err != nil {
		return err
	}
	if prepStmts.get, err = myDB.Prepare(`
			SELECT *
			FROM subscriptions
	    WHERE id = $1
	  `); err != nil {
		return err
	}
	if prepStmts.list, err = myDB.Prepare(`
			SELECT *
			FROM subscriptions
  	`); err != nil {
		return err
	}

	return err
}

package subscription

import (
	"database/sql"
	"errors"
)

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

	if _, err := myDB.Exec(sqlSchema); err != nil {
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

	return subscription, errors.New("No subscription found")
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

	if prepStmts.insert, err = myDB.Prepare(sqlInsert); err != nil {
		return err
	}
	if prepStmts.destroy, err = myDB.Prepare(sqlDestroy); err != nil {
		return err
	}
	if prepStmts.get, err = myDB.Prepare(sqlRetrieve); err != nil {
		return err
	}
	if prepStmts.list, err = myDB.Prepare(sqlList); err != nil {
		return err
	}

	return err
}

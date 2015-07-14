package subscription

import "database/sql"

type Persistance struct{}
type Persist interface {
	get(id string) (Subscription, error)
	list() ([]Subscription, error)
	destroy(id string) (Subscription, error)
}

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

func (p Persistance) get(id string) (Subscription, error) {
	var sub Subscription

	err := prepStmts.get.QueryRow(id).Scan(&sub.ID, &sub.EventType, &sub.Context, &sub.AccountID, &sub.CreatedAt)

	if err != nil {
		return Subscription{}, err
	}

	return sub, err
}

func (p Persistance) list() ([]Subscription, error) {
	rows, err := prepStmts.list.Query()

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	subs := []Subscription{}

	for rows.Next() {
		var sub Subscription

		if err := scan(rows, &sub); err != nil {
			return nil, err
		}

		subs = append(subs, sub)
	}

	err = rows.Err()
	return subs, err
}

func (sub *Subscription) insert() error {
	err := prepStmts.insert.QueryRow(sub.EventType, sub.Context, sub.AccountID.String()).Scan(&sub.ID, &sub.CreatedAt)

	return err
}

func (p Persistance) destroy(id string) (Subscription, error) {
	sub, err := p.get(id)
	if err != nil {
		return sub, err
	}

	_, err = prepStmts.destroy.Query(id)
	return sub, err
}

func scan(rows *sql.Rows, sub *Subscription) error {
	return rows.Scan(&sub.ID, &sub.EventType, &sub.Context, &sub.AccountID, &sub.CreatedAt)
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

package subscription

const sqlSchema = `
		CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
		CREATE TABLE IF NOT EXISTS subscriptions (
			id UUID PRIMARY KEY,
			event_type VARCHAR(64),
			context VARCHAR(64),
			account_id UUID,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP )
	`
const sqlInsert = `
		INSERT INTO subscriptions
		( id,
			event_type,
			context,
			account_id
		)
		VALUES ( uuid_generate_v4(), $1, $2, $3 )
		RETURNING id, created_at
	`
const sqlDestroy = `
		DELETE
		FROM subscriptions
		WHERE id = $1
	`
const sqlRetrieve = `
		SELECT *
		FROM subscriptions
		WHERE id = $1
	`
const sqlList = `
		SELECT *
		FROM subscriptions
	`

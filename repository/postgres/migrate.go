package postgres

import (
	"context"
	"database/sql"
)

func CreateTableForUsers(ctx context.Context, db *sql.DB) error {
	queru := `
CREATE TABLE IF NOT EXISTS users(
name TEXT PRIMARY KEY,
status BOOLEAN NOT NULL
);
`

	_, err := db.ExecContext(ctx, queru)
	return err
}

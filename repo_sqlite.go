package users

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const driverNameSQLite string = "sqlite3"

// openSQlite opens SQLite database using specified path
// and executes create table statement.
// CREATE TABLE IF NOT EXISTS users (
// 	username text primary key,
// 	pwdhash text not null,
// 	enabled bool default true,
// 	attrs blob,
// 	created datetime,
// 	updated datetime
// )
// If table named "users"
// already exists in the database the create statement will have
// no effect and will return nil error.
func openSQlite(path string) (*sql.DB, error) {
	db, err := sql.Open(driverNameSQLite, path)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(1)
	_, err = db.Exec(statementSQLiteCreateTableUsers)
	return db, err
}

package users

import (
	"database/sql"

	_ "github.com/jackc/pgx/v4/stdlib"
)

const driverNamePG string = "pgx"

// openPG opens connects to Postgres database using specified dsn
// and executes create table statement.
// CREATE TABLE IF NOT EXISTS users (
// 	username varchar(90) primary key,
// 	pwdhash varchar(150) not null,
// 	enabled bool default true,
// 	attrs bytea,
// 	created timestamp with time zone not null default now(),
// 	updated timestamp with time zone not null default now()
// )
// If table named "users"
// already exists in the database the create statement will have
// no effect and will return nil error.
func openPG(dsn string) (*sql.DB, error) {
	db, err := sql.Open(driverNamePG, dsn)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(statementPostgresCreateTableUsers)
	return db, err
}

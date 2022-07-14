package users

import (
	"database/sql"
)

// NewNewWithPG return Repository and Authenticator
// using connection to Postgres DB via specified DSN under the hood.
func NewWithPG(dsn string) (Repository, Authenticator, error) {
	var (
		db   *sql.DB
		repo Repository
		auth Authenticator
		err  error
	)

	db, err = openPG(dsn)

	repo = newSQLRepo(db)
	auth = newBcryptAuth(repo)

	return repo, auth, err
}

// NewWithSQLite return Repository and Authenticator
// using SQLite DB with specified path under the hood.
func NewWithSQLite(path string) (Repository, Authenticator, error) {
	var (
		db   *sql.DB
		repo Repository
		auth Authenticator
		err  error
	)

	db, err = openSQlite(path)

	repo = newSQLRepo(db)
	auth = newBcryptAuth(repo)
	return repo, auth, err
}

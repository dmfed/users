package users

// creates table users in Postgres db
var statementPostgresCreateTableUsers = `
	CREATE TABLE IF NOT EXISTS users (
    	username varchar(90) primary key,
    	secret varchar(150) not null,
    	enabled bool default true,
    	attrs bytea,
    	created timestamp with time zone not null default now(),
		updated timestamp with time zone not null default now()
	)`

// creates table users in SQLite db
var statementSQLiteCreateTableUsers = `
	CREATE TABLE IF NOT EXISTS users (
		username text primary key, 
		secret text not null, 
		enabled bool, 
		attrs blob, 
		created datetime, 
		updated datetime
	)`

var (
	statementGet = `
	SELECT username, secret, enabled, attrs, created, updated 
	FROM users
	WHERE username = $1`

	statementPut = `
	INSERT INTO users (username, secret, enabled, attrs, created, updated)
	VALUES ($1, $2, $3, $4, $5, $6)`

	statementUpd = `
	UPDATE users SET
	secret = $1,
	enabled = $2,
	attrs = $3,
	updated = $4
	WHERE username = $5`

	statementRenameUsr = `
	UPDATE users SET
	username = $1,
	updated = $2
	WHERE username = $3`

	statementDel = `
	DELETE FROM users
	WHERE username = $1`
)

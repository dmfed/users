package users

import (
	"database/sql"
	"encoding/json"
	"errors"
	"time"
)

func getUsr(db *sql.DB, name string) (User, error) {
	row := db.QueryRow(statementGet, name)
	var (
		u User
		b []byte
	)
	err := row.Scan(&u.Username, &u.Secret, &u.Enabled, &b, &u.Created, &u.Updated)
	u.Attrs = bytesToMap(b)
	return u, replaceNoRowsErr(err)
}

func putUsr(db *sql.DB, u *User) error {
	_, err := db.Exec(statementPut, u.Username, u.Secret, u.Enabled, mapToBytes(u.Attrs), u.Created, u.Updated)
	return err
}

func updUsr(db *sql.DB, u *User) error {
	return checkErr(db.Exec(statementUpd, u.Secret, u.Enabled, mapToBytes(u.Attrs), u.Updated, u.Username))
}

func renameUsr(db *sql.DB, o, n string, t time.Time) error {
	return checkErr(db.Exec(statementRenameUsr, n, t, o))
}

func delUsr(db *sql.DB, u string) error {
	return checkErr(db.Exec(statementDel, u))
}

func mapToBytes(m map[string]string) (b []byte) {
	b, _ = json.Marshal(m)
	return
}

func bytesToMap(b []byte) (m map[string]string) {
	json.Unmarshal(b, &m)
	return
}

func checkErr(res sql.Result, err error) error {
	if res == nil {
		return err
	}
	if n, _ := res.RowsAffected(); n == 0 && err == nil {
		err = ErrNotFound
	}
	return err
}

func replaceNoRowsErr(err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		err = ErrNotFound
	}
	return err
}

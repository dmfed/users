package users

import (
	"database/sql"
	"time"
)

type userRepoSQL struct {
	db *sql.DB
}

// NewSQLRepo accepts SQL database instance and return Repostory.
// Note that Put, Upd and Rename methods enforce non-zero timestapms
// of User struct. If either Created or Updated fields are zero time
// values (IsZero() returns true), then they are updated with time.Now()
// Put method also sets Enabled field to true for User struct.
// This way it is safe to actually fill only Name and Secret fields
// of User struct and pass it to Put method.
func newSQLRepo(db *sql.DB) Repository {
	return &userRepoSQL{db}
}

func (u *userRepoSQL) Get(username string) (User, error) {
	return getUsr(u.db, username)
}

func (u *userRepoSQL) Put(user User) error {
	if !isValidUser(&user) {
		return ErrValidateUser
	}
	checkAndCorrectTimestamps(&user)
	user.Enabled = true
	return putUsr(u.db, &user)
}

func (u *userRepoSQL) Upd(user User) error {
	if !isValidUser(&user) {
		return ErrValidateUser
	}
	checkAndCorrectTimestamps(&user)
	return updUsr(u.db, &user)
}

func (u *userRepoSQL) Del(username string) error {
	return delUsr(u.db, username)
}

func (u *userRepoSQL) Rename(old, new string) error {
	return renameUsr(u.db, old, new, time.Now())
}

func (u *userRepoSQL) Close() error {
	return u.db.Close()
}

// a bit of enforcement of data integrity
func checkAndCorrectTimestamps(u *User) {
	now := time.Now()
	if u.Created.IsZero() {
		u.Created = now
	}
	if u.Updated.IsZero() {
		u.Updated = now
	}
}

// thou shalt not ever omit these fields
func isValidUser(u *User) bool {
	return u.Username != "" && u.Secret != ""
}

## Package users

**users** is very basic CRUD for storing user data including username, secret,
attributes (Go map[string]string data type) and created/updated timestamps. The package is
intended to be used in smaller projects requiring user authentication. Think of
your pet project involving access by multiple users. 

Currently two underlying storage types are supported: Postgres and SQLite.

### Usage

User info is stored in User struct:

```go
type User struct {
	Username string            // username
	Secret   string            // secret (e.g. password hash)
	Enabled  bool              // is user active
	Attrs    map[string]string // various stuff you might put here
	Created  time.Time         // create time
	Updated  time.Time         // last update time
}
```

All the useful stuff can be done with two interfaces exposed by the package.

```go
type Authenticator interface {
	Authenticate(username, password string) (bool, error) // Authenticate returns true if password for username checks out.
}

// Repository is a basic CRUD interface to
// store User structs.
type Repository interface {
	Get(username string) (User, error)           // Get returns user with username
	Put(user User) error                         // Put saves User struct
	Upd(user User) error                         // Update updates User
	Del(username string) error                   // Del deletes user with username
	Rename(oldname string, newname string) error // Rename renames user with oldname setting username to newname
	Close() error                                // Closes safely closes the repository
}
```

Usage example can be found in **example** directory. 

When instantiated with **New** function the Authenticator uses bcrypt unde the hood
to compare provided password with stored hash. 

The package contains a function **BcryptHash** to generate hash of a password (a string).

```go
package main

import (
	"fmt"

	"github.com/dmfed/users"
)

func main() {
	repo, auth, err := users.NewWithSQLite("./testfoo.db")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = repo.Put(users.User{
		Username: "testuser",
		Secret:   users.BcryptHash("testsecret"),
		Attrs: map[string]string{
			"happy": "yes",
		},
	})

	if err != nil {
		fmt.Println("OMG we failed", err)
	}

	u, _ := repo.Get("testuser")
	fmt.Println(u)

	ok, _ := auth.Authenticate("testuser", "testsecret") // check that password for testuser is "testsecret"
	// ok is true, err is nil
	if !ok {
		fmt.Println("password does not match username 'testuser'")
	}

	repo.Del("testuser")
}
```
**Put** and **Upd** methods enforce non-zero timestamps for User struct. Actually the only required values for User struct
are Username and Secret. Failing to fill these will result in non-nil errors from Put and Upd.

Also note that **Put** method sets user **Enabled** field to **true**.

### command-line utility

**cmd** directory contains a handy cli utility to manage users. 
```bash
> cd cmd
> go build
> ./cmd --help                                                                                                                                   
Usage of ./cmd:
  -a	<username> <hash> add user
  -c	<username> <password> check user password against db record
  -d	<username> delete user with username
  -h	<password> print hash of supplied password
  -s string
    	data source (sqlite filename or Postgres DSN formatted as postgres://username:password@localhost:5432/db?sslmode=disable)
  -show
    	<username> show database record a user
  -t string
    	data source type (sqlite or postgres) (default "sqlite")
  -toggle
    	toggle user "enabled" property
  -u	<username> <hash> update user record with provided hash

> ./cmd -s testdb.db -a testuser $(./cmd -h password)                                                                                              ✹main 
added user testuser

> ./cmd -s testdb.db -show testuser                                                                                                              ✹ ✭main 
Name: testuser
Hash: $2a$10$kwINOS5wz4L6Ozrho9yEFuRqSo40tMxAXnz0ssC7TThFW.LDkvD.2
Enabled: true
Attrs: map[]
Created: 2022-07-14 20:37:15 +0300 +0300
Updated: 2022-07-14 20:37:15 +0300 +0300

>    
```
### Thanks

Thanks to **mattn** the creator of excellent library https://github.com/mattn/go-sqlite3

Thanks to **jackc** the creator of pgx https://github.com/jackc/pgx
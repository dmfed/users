package main

import (
	"flag"
	"fmt"

	"github.com/dmfed/users"
)

func getHash(password string) {
	fmt.Println(users.BcryptHash(password))
}

func getUsr(r users.Repository, username string) {
	u, err := r.Get(username)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(u)
}

func addUsr(r users.Repository, username, hash string) {
	u := users.User{
		Username: username,
		Secret:   hash,
	}

	err := r.Put(u)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("added user %s\n", username)
}

func updUserHash(r users.Repository, username, hash string) {
	u, err := r.Get(username)
	if err != nil {
		fmt.Println(err)
		return
	}
	u.Secret = hash
	err = r.Upd(u)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("updated user %s\n", username)
}

func checkPwd(a users.Authenticator, username, password string) {
	ok, err := a.Authenticate(username, password)
	if err != nil {
		fmt.Println(err)
		return
	}
	if ok {
		fmt.Printf("correct password for user %s\n", username)
	} else {
		fmt.Printf("incorrect password for user %s\n", username)
	}
}

func toggleUsr(r users.Repository, username string) {
	u, err := r.Get(username)
	if err != nil {
		fmt.Println(err)
		return
	}
	u.Enabled = !u.Enabled
	err = r.Upd(u)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("enabled set to %v for user %s\n", u.Enabled, username)
}

func delUsr(r users.Repository, username string) {
	err := r.Del(username)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("deleted user %s\n", username)
}

func main() {
	var (
		add  bool
		chk  bool
		upd  bool
		tog  bool
		del  bool
		show bool
		hash bool

		sourcetype string
		source     string
	)

	flag.BoolVar(&upd, "u", false, "<username> <hash> update user record with provided hash")
	flag.BoolVar(&add, "a", false, "<username> <hash> add user")
	flag.BoolVar(&chk, "c", false, "<username> <password> check user password against db record")
	flag.BoolVar(&tog, "toggle", false, "toggle user \"enabled\" property")
	flag.BoolVar(&show, "show", false, "<username> show database record a user")
	flag.BoolVar(&hash, "h", false, "<password> print hash of supplied password")
	flag.BoolVar(&del, "d", false, "<username> delete user with username")
	flag.StringVar(&sourcetype, "t", "sqlite", "storage type: sqlite or postgres")
	flag.StringVar(&source, "s", "", "data source (sqlite filename or Postgres DSN formatted as postgres://username:password@localhost:5432/db?sslmode=disable)")

	flag.Parse()

	if hash {
		getHash(flag.Arg(0))
		return
	}

	var (
		repo users.Repository
		auth users.Authenticator
		err  error
	)

	switch sourcetype {
	case "sqlite":
		repo, auth, err = users.NewWithSQLite(source)
	case "postgres":
		repo, auth, err = users.NewWithPG(source)
	default:
		err = fmt.Errorf("unknown storage type")
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	if show {
		getUsr(repo, flag.Arg(0))
	} else if chk {
		checkPwd(auth, flag.Arg(0), flag.Arg(1))
	} else if add {
		addUsr(repo, flag.Arg(0), flag.Arg(1))
	} else if upd {
		updUserHash(repo, flag.Arg(0), flag.Arg(1))
	} else if tog {
		toggleUsr(repo, flag.Arg(0))
	} else if del {
		delUsr(repo, flag.Arg(0))
	}
}

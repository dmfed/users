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

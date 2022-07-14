package users

import (
	"os"
	"testing"
)

var dbname = "./testfoo.db"

func TestPGRepo(t *testing.T) {
	dsn := os.Getenv("TESTDB")
	if dsn == "" {
		t.Skip("TESDB env not set skipping Postgres test")
	}
	db, err := openPG(dsn)
	if err != nil {
		t.Error(err)
		return
	}
	testRepo(t, newSQLRepo(db))
}

func TestSQLiteRepo(t *testing.T) {
	os.Remove(dbname)
	db, err := openSQlite(dbname)
	if err != nil {
		t.Error(err)
		return
	}
	testRepo(t, newSQLRepo(db))
	os.Remove(dbname)
}

func TestNewWithSQLite(t *testing.T) {
	os.Remove(dbname)
	if _, _, err := NewWithSQLite(dbname); err != nil {
		t.Error(err)
	}
	os.Remove(dbname)
}

func TestNewWithPG(t *testing.T) {
	dsn := os.Getenv("TESTDB")
	if dsn == "" {
		t.Skip("TESTDB env not set skipping New test for Postgres")
	}
	if _, _, err := NewWithPG(dsn); err != nil {
		t.Error(err)
	}
}

func testRepo(t *testing.T, r Repository) {
	defer r.Close()

	// test Get yields error on random query
	_, err := r.Get("obi-wan")
	if err == nil {
		t.Errorf("Get does not fail on random query")
	}

	// testing Put fails on mepty user struct
	err = r.Put(User{})
	if err == nil {
		t.Errorf("repo Puts empty user")
	}

	u := User{
		Username: "testuser",
		Secret:   BcryptHash("testpassword"), // using hash here to check Authenticate later
		Enabled:  true,
		Attrs: map[string]string{
			"attr1": "hello",
		},
	}

	// testing we can put and get back
	err = r.Put(u)
	if err != nil {
		t.Error(err)
	}

	u2, err := r.Get("testuser")
	if err != nil {
		t.Error(err)
	}

	if u2.Username != u.Username || u.Secret != u2.Secret ||
		u2.Enabled != u.Enabled {
		t.Errorf("got something useless from repo")
	}

	for k, v := range u.Attrs {
		v2, ok := u2.Attrs[k]
		if v2 != v || !ok {
			t.Errorf("Attrs maps are not equal")
		}
	}

	u2.Attrs["added"] = "newvalue"
	u2.Enabled = false

	// testing we can update
	if err = r.Upd(u2); err != nil {
		t.Error(err)
	}

	u2, err = r.Get("testuser")
	if err != nil {
		t.Error(err)
	}

	if v, ok := u2.Attrs["added"]; v != "newvalue" || !ok {
		t.Errorf("attributes did not update")
	}

	if u2.Enabled {
		t.Errorf("enabled did not update")
	}

	// test Update fails on invalid username
	err = r.Upd(User{Username: "randomname", Secret: "randomhash"})
	if err == nil {
		t.Errorf("Upd does not fail on random username")
	}

	// test Upd fails on empty user struct
	err = r.Upd(User{})
	if err == nil {
		t.Errorf("Upd does not fail on empty User struct")
	}

	// testing rename works
	if err = r.Rename("testuser", "newtestuser"); err != nil {
		t.Error(err)
	}

	u3, err := r.Get("newtestuser")
	if err != nil {
		t.Error(err)
	}

	if u3.Enabled != u2.Enabled || u3.Secret != u2.Secret {
		t.Errorf("rename didn't work")
	}

	auth := newBcryptAuth(r)

	// testing Authenticate works
	if ok, err := auth.Authenticate("newtestuser", "testpassword"); !ok || err != nil {
		t.Errorf("Authenticate fails on valid case")
	} else if ok, err := auth.Authenticate("newtestuser", "randomstuff"); ok || err == nil {
		t.Errorf("Authenticate doesn't fail on invalid case")
	}

	// testing del works
	if err = r.Del("newtestuser"); err != nil {
		t.Error(err)
	}
}

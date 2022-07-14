package users

import (
	"fmt"
	"time"
)

// User is a very basic struct representing user
type User struct {
	Username string            // username
	Secret   string            // secret (e.g. password hash)
	Enabled  bool              // is user active
	Attrs    map[string]string // various stuff you might put here
	Created  time.Time         // create time
	Updated  time.Time         // last update time
}

func (u User) String() string {
	return fmt.Sprintf(
		"Name: %s\nHash: %s\nEnabled: %v\nAttrs: %s\nCreated: %s\nUpdated: %s\n",
		u.Username, u.Secret, u.Enabled, u.Attrs, u.Created.Truncate(time.Second), u.Updated.Truncate(time.Second))
}

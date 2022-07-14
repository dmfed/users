package users

type Authenticator interface {
	Authenticate(username, password string) (bool, error) // Authenticate returns true if password for username checks out. If error is not nil the boolean resualt is unusable.
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

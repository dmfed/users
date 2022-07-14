package users

import (
	"golang.org/x/crypto/bcrypt"
)

type bcryptAuth struct {
	repo Repository
}

// NewBcryptAuth returns Authenticator using bcrypt
// algorithm in Authenticate method to compare provided
// password with stored hash.
func newBcryptAuth(repo Repository) Authenticator {
	return &bcryptAuth{repo}
}

// Authenticate implements Authenticator interface.
func (b *bcryptAuth) Authenticate(username, password string) (bool, error) {
	u, err := b.repo.Get(username)
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Secret), []byte(password))
	if err != nil {
		err = ErrIncorrectPassword
	}
	return (err == nil), err
}

// BcryptHash return bcrypt hash of provided password.
func BcryptHash(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}

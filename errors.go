package users

import "errors"

var (
	ErrNotFound           = errors.New("no such user")                                     // ErrNotFound is returned when trying to do something with uknown username
	ErrValidateUser       = errors.New("missing username or secret field for user")        // ErrValidateUser is returned when User struct is missing Username or Secret fields
	ErrIncorrectPassword  = errors.New("incorrect password")                               // ErrIncorrectPassword is returned when user password does not check out
	ErrUnknownStorageType = errors.New("unknown storage type, use 'postgres' or 'sqlite'") // ErrUnknownStorageType is returned when unknown value is passed as storage type in RepoConfig
)

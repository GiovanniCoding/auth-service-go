package errors

import "errors"

var (
	ErrInvalidToken      = errors.New("invalid token")
	ErrUserPassInvalid   = errors.New("user or password incorrect")
	ErrInternalServerErr = errors.New("internal server error")
	ErrDBConnection      = errors.New("error connecting to database")
	ErrUserAlreadyExists = errors.New("user already exists")
)

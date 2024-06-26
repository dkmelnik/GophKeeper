package errs

import "errors"

var (
	ErrNotFound           = errors.New("not found")
	ErrFailedCreate       = errors.New("failed to create")
	ErrIsExist            = errors.New("is exist")
	ErrFailedUpdate       = errors.New("failed to update")
	ErrInvalidToken       = errors.New("invalid token")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

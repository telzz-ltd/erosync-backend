package users

import "errors"

var (
	ErrEmailExist         = errors.New("email already exist")
	ErrInvalidCredentials = errors.New("incorrect email or password")
)

package repository

import "errors"

var (
	ErrInvalidArgument = errors.New("invalid argument")
	ErrAlreadyExist    = errors.New("already exist")
	ErrNotFound        = errors.New("not found")
)

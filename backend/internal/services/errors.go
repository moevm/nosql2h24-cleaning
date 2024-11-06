package services

import "errors"

var (
	// Auth errors
	ErrHashing           = errors.New("hashing error occurred")
	ErrIncorrectPassword = errors.New("incorrect password")

	// User errors
	ErrUserAlreadyExist = errors.New("user already exist")
	ErrUserNotFound     = errors.New("user not found")

	// Orders errors
	// Service errors
)

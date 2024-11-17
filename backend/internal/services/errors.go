package services

import "errors"

var (
	// Auth errors
	ErrHashing           = errors.New("hashing error occurred")
	ErrIncorrectPassword = errors.New("incorrect password")

	// User errors
	ErrUserAlreadyExist = errors.New("user already exist")
	ErrUserNotFound     = errors.New("user not found")
	ErrAddressNotFound  = errors.New("address not found")
	// Orders errors
	ErrOrderAlreadyExist = errors.New("order already exist")
	ErrOrderNotFound     = errors.New("order not found")
	// Service errors
	ErrServiceAlreadyExist = errors.New("service already exist")
	ErrServiceNotFound     = errors.New("service not found")
)

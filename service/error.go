package service

import "errors"

var (
	//ErrInternalServerError Error caused by system error
	ErrInternalServerError = errors.New("internal server error")

	//ErrNotFound Error when item is not found
	ErrNotFound = errors.New("data was not found")

	//ErrInvalidSpec Error when data given is not valid on update or insert
	ErrInvalidData = errors.New("given data is not valid")

	//ErrLogin Error when email or password is incorrect
	ErrLogin = errors.New("email or password is incorrect")

	//ErrRegister Error when duplicate entry on email or phone_number
	ErrRegister = errors.New("email or phone_number already registered")
)

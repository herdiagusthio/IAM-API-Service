package user

import (
	"time"
)

type User struct {
	ID           uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
	Name         string
	Email        string
	Phone_number string
	Password     string
	RoleID       uint
}

func NewUserFromHandler(
	name string,
	email string,
	phoneNumber string,
	password string,
	createdAt time.Time,
	updatedAt time.Time,
) User {
	return User{
		ID:           0,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		DeletedAt:    nil,
		Name:         name,
		Email:        email,
		Phone_number: phoneNumber,
		Password:     password,
		RoleID:       2,
	}
}

func NewAdminFromHandler(
	name string,
	email string,
	phoneNumber string,
	password string,
	createdAt time.Time,
	updatedAt time.Time,
) User {
	return User{
		ID:           0,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		DeletedAt:    nil,
		Name:         name,
		Email:        email,
		Phone_number: phoneNumber,
		Password:     password,
		RoleID:       1,
	}
}

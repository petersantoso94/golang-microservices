package user

import (
	"errors"

	"gorm.io/gorm"
)

// ErrNotFound signifies that a single requested object was not found.
var ErrNotFound = errors.New("not found")

// User is a user business object.
type User struct {
	ID   int64 `gorm:"primarykey"`
	Name string
	gorm.Model
}

// Service defines the interface exposed by this package.
type Service interface {
	CreateUser(User) (int64, error)
	GetUser(id int64) (User, error)
	GetUsers() ([]User, error)
}

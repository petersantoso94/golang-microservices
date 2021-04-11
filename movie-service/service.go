package movie

import (
	"errors"

	userSvc "github.com/petersantoso94/golang-microservices/user-service"
)

// ErrNotFound signifies that a single requested object was not found.
var ErrNotFound = errors.New("not found")

// User is a user business object.
type Movie struct {
	ID    int64 `gorm:"primarykey"`
	Name  string
	Owner userSvc.User
}

// Service defines the interface exposed by this package.
type Service interface {
	GetMovies() ([]*Movie, error)
	GetUserMovie(userId int64) ([]*Movie, error)
	CreateMovie(Movie) error
}

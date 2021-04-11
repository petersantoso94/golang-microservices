package movie

import (
	"errors"
)

// ErrNotFound signifies that a single requested object was not found.
var ErrNotFound = errors.New("not found")

// User is a user business object.
type Movie struct {
	ID      int64 `gorm:"primarykey"`
	Name    string
	OwnerID int64
}

// Service defines the interface exposed by this package.
type Service interface {
	GetMovies() ([]*Movie, error)
	GetUserMovie(OwnerID int64) ([]*Movie, error)
	CreateMovie(Movie) error
}

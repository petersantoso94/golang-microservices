package movie

import (
	"errors"

	userSvc "github.com/petersantoso94/golang-microservices/user-service"
)

// ErrNotFound signifies that a single requested object was not found.
var ErrNotFound = errors.New("not found")

// User is a user business object.
type Movie struct {
	ID    int64
	Name  string
	Owner userSvc.User
}

// Service defines the interface exposed by this package.
type Service interface {
	GetMovie(name string) (Movie, error)
	GetMovies() (map[int64]Movie, error)
	GetUserMovie(userId []int64) (map[int64]Movie, error)
}

package core

import (
	movieSvc "github.com/petersantoso94/golang-microservices/movie-service"
	"github.com/petersantoso94/golang-microservices/user-service"
)

type service struct {
	// a database dependency would go here but instead we're going to have a static map
	m map[int64]movieSvc.Movie
	movieSvc.Service
}

// NewService instantiates a new Service.
func NewService( /* a database connection would be injected here */ ) movieSvc.Service {
	return &service{
		m: map[int64]movieSvc.Movie{
			1: {ID: 1, Name: "Ironman", Owner: user.User{ID: 1}},
			2: {ID: 2, Name: "Shrek", Owner: user.User{ID: 2}},
			3: {ID: 3, Name: "SpongeBob", Owner: user.User{ID: 2}},
		},
	}
}

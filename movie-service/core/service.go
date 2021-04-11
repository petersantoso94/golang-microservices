package core

import (
	movieSvc "github.com/petersantoso94/golang-microservices/movie-service"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type service struct {
	db *gorm.DB
}

// NewService instantiates a new Service.
func NewService() movieSvc.Service {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	db.AutoMigrate(&movieSvc.Movie{})
	return &service{
		db: db,
	}
}

func (s *service) GetMovies() (result []*movieSvc.Movie, err error) {
	if ok := s.db.Find(&result); ok.RowsAffected == 0 || ok.Error != nil {
		return nil, ok.Error
	}
	return result, nil
}

func (s *service) GetUserMovie(userId int64) (result []*movieSvc.Movie, err error) {
	if ok := s.db.Where(&movieSvc.Movie{OwnerID: userId}).Find(&result); ok.RowsAffected == 0 || ok.Error != nil {
		return nil, ok.Error
	}
	return result, nil
}

func (s *service) CreateMovie(movie movieSvc.Movie) error {
	if result := s.db.Create(&movie); result.RowsAffected == 0 || result.Error != nil {
		return result.Error
	}
	return nil
}

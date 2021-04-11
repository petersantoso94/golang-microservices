package core

import (
	userSvc "github.com/petersantoso94/golang-microservices/user-service"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type service struct {
	db *gorm.DB
}

// NewService instantiates a new Service.
func NewService() userSvc.Service {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	db.AutoMigrate(&userSvc.User{})
	return &service{
		db: db,
	}
}

func (s *service) GetUser(id int64) (result *userSvc.User, err error) {
	if ok := s.db.First(&result, id); ok.RowsAffected > 0 {
		return result, nil
	}
	return result, userSvc.ErrNotFound
}

func (s *service) GetUsers() (result []*userSvc.User, err error) {
	if ok := s.db.Find(&result); ok.RowsAffected == 0 || ok.Error != nil {
		return nil, ok.Error
	}
	return
}

func (s *service) CreateUser(user userSvc.User) (id int64, err error) {
	if ok := s.db.Create(&user); ok.RowsAffected == 0 || ok.Error != nil {
		return 0, ok.Error
	}
	return user.ID, nil
}

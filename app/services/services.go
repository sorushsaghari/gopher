package services

import (
	"../models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)


type Service struct {
	DB *gorm.DB
}


func WithGorm(dialect string, connectionInfo string) (*Service, error) {
		db, err := gorm.Open(dialect, connectionInfo)
		if err != nil {
			return nil, err
		}

	return &Service{DB: db}, nil
}

func (s *Service) Migrate() error{
	return 	s.DB.AutoMigrate(&models.User{}).Error
}



func (s *Service) Close() error {
	return s.DB.Close()
}

package services

import (
	"../models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)
type ServeConfig func(s *Service) error

type Service struct {
	service
	User UserService
	Nurse NurseService
	Chat MessageService
}

type service struct {
	db *gorm.DB
}

func NewService(cfgs ...ServeConfig) (*Service, error) {
	var s Service
	for _, cfg := range cfgs {
		if err := cfg(&s); err != nil {
			return nil, err
		}
	}
	return &s, nil
}


func WithGorm(dialect string, connectionInfo string) ServeConfig {
	return func(s *Service) error {
		db, err := gorm.Open(dialect, connectionInfo)
		if err != nil {
			return err
		}
		s.db = db
		return nil
	}
}

func WithUser() ServeConfig {
	return func(s *Service) error {
		s.User = NewUserService(s.db)
		return nil
	}
}

func WithChat() ServeConfig {
	return func(s *Service) error {
		s.Chat = NewMessageService(s.db)
		return nil
	}
}

func (s *Service) Migrate() error{
	return 	s.db.AutoMigrate(&models.User{}).Error
}



func (s *Service) Close() error {
	return s.db.Close()
}

package models

import "github.com/jinzhu/gorm"

type ServicesConfig func(s *Service) error

type Service struct {
	db   *gorm.DB
}


func NewService(cfgs ...ServicesConfig) (*Service, error) {
	var s Service
	for _, cfg := range cfgs {
		if err := cfg(&s); err != nil {
			return nil, err
		}
	}
	return &s, nil
}


func WithGorm(dialect string, connectionInfo string) ServicesConfig{
	return func(s *Service) error {
		db, err := gorm.Open(dialect, connectionInfo)
		if err != nil {
			return err
		}
		s.db = db
		return nil
	}
}


func (s *Service) Close() error {
	return s.db.Close()
}

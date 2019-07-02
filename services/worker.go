package services

import (
	"../models"
	"github.com/jinzhu/gorm"
)

type NurseDto struct {
	UserDto
	PersonalId string `json:"personal_id" valid:"-"`
	Licence    string `json:"licence" valid:"-"`
}

type NurseDb interface {
	FindById(id uint) (models.Nurse, error)
	Insert(NurseDto) error
	All() (*[] models.Nurse, error)
}

type NurseService interface {
	NurseDb
}

type nurseService struct {
	service
	NurseService
}

func (n nurseService) NewNurseService(db *gorm.DB) *nurseService {
	service := service{db}
	return &nurseService{service: service}
}

func (n nurseService) Insert(object NurseDto) error {
	err := userService.Insert(object.UserDto)
}
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
	Service
}

func NewNurseService(db *gorm.DB) *nurseService {
	service := Service{DB: db}
	return &nurseService{ Service: service}
}

func (n nurseService) NewNurseService(db *gorm.DB) *nurseService {
	service := Service{db}
	return &nurseService{ service}
}

func (n nurseService) Insert(object NurseDto) error {
}
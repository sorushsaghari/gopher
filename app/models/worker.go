package models

import "github.com/jinzhu/gorm"

type Nurse struct {
	gorm.Model
	User User `gorm:"polymorphic:Owner;"`
	PersonalId string
	Licence string
}
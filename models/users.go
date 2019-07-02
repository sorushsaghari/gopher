package models

import (
	"github.com/jinzhu/gorm"
)


type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `gorm:"not null;unique_index" json:"email"`
	Password string `json:"password"`
	OwnerID   int
	OwnerType string
}

func (u User) IsNurse() bool {
	if u.OwnerType == "Nurse"{
		return true
	}
	return false
}

func (u User) IsCustomer() bool {
	if u.OwnerType == "Customer"{
		return true
	}
	return false
}
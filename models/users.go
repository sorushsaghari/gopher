package models

import (
	"github.com/jinzhu/gorm"
)


type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `gorm:"not null;unique_index" json:"email"`
	Password string `json:"password"`
}



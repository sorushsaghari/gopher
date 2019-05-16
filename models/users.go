package models

import "github.com/jinzhu/gorm"

type UserDb interface {
}

type UserService interface {
	UserDb
	Authenticate(password, email string) (bool, error)
}
type userService struct {
	db *gorm.DB
	UserService
}
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `gorm:"not null;unique_index" json:"email"`
	Password string `json:"password"`
}

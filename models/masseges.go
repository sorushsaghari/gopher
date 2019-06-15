package models

import "github.com/jinzhu/gorm"

type Message struct {
	gorm.Model
	Text string `json:"name"`
	UserID int
	User User
}
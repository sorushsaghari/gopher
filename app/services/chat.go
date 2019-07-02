package services

import "github.com/jinzhu/gorm"

type MessageDto struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

type MessageDB interface {

}

type MessageService interface {
	MessageDB

}
type messageService struct {
	db *gorm.DB
	MessageService
}

func NewMessageService(db *gorm.DB) *messageService {
	return &messageService{db: db}
}
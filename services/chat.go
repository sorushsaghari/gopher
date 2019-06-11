package services

import "github.com/jinzhu/gorm"

type MessageDto struct {
	User UserDto `json:"user"`
	Text string  `json:"text"`
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
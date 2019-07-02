package app

import (
	"./services"
)

type AppConfig func(s *App) error

type App struct {
	User services.UserService
	Nurse services.NurseService
	Chat services.MessageService
}

func NewApp(cfgs ...AppConfig) (*App, error) {
	var s App
	for _, cfg := range cfgs {
		if err := cfg(&s); err != nil {
			return nil, err
		}
	}
	return &s, nil
}


func WithUser(s services.Service) AppConfig {
	return func(a *App) error {
		a.User = services.NewUserService(s.DB)
		return nil
	}
}

func WithChat(s services.Service) AppConfig {
	return func(a *App) error {
		a.Chat = services.NewMessageService(s.DB)
		return nil
	}
}

func WithNurse(s services.Service) AppConfig {
	return func(a *App) error {
		a.Chat = services.NewNurseService(s.DB)
		return nil
	}
}

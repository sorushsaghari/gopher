package controllers

import (
	"../models"
)

type UserController struct {
	us models.UserService
}

func NewUserController(us models.UserService) *UserController{
	return  &UserController{
		us: us,
	}
}


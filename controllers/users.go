package controllers

import (
	"../models"
	"github.com/gin-gonic/gin"
	"time"
	"github.com/dgrijalva/jwt-go"
	"../config"
	"../services"
	"net/http"
	"fmt"
)

type UserController struct {
	us services.UserService
}

func NewUserController(us services.UserService) *UserController{
	return  &UserController{
		us: us,
	}
}

func (uc* UserController) SignIn(c* gin.Context)  {

	var cred config.Credentials
	c.ShouldBindJSON(&cred)
	authenticated, err:= uc.us.Authenticate(cred.Email, cred.Password)
	if err != nil {
		c.JSON(404, map[string]string{"error": err.Error()})
		return
	}

	if !authenticated {
		c.JSON(404, map[string]string{"error": models.ErrNotFound.Error()})
		return
	}

	expirationTime := time.Now().Add(60 * time.Minute)
	claims := &config.JWT{
		Email: cred.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(config.JwtKey)
	if err != nil {
		c.JSON(500, map[string]string{"error": err.Error()})
		return
	}

	c.JSON(200, map[string]string{"token": tokenString})
}

func (uc* UserController) Me(c* gin.Context)  {
	//me, _ := uc.us.FindById(1 )
	claim, err := config.ParseToken( c.GetHeader("token"))
	if err != nil {
		c.JSON(404, err)
		return
	}
	me, err := uc.us.FindByEmail(claim.Email)
	if err != nil {
		c.JSON(404, err)
		return
	}
	c.JSON(200, me)
}



func (uc* UserController) Create(c* gin.Context) {
	var json services.UserDto
	c.ShouldBindJSON(&json)
	result, err := uc.us.Validate(json)
	if err != nil {
		fmt.Println(err)
		 c.JSON(http.StatusBadRequest, gin.H{"err":err.Error()})
		 return
	}
	if !result {
		c.JSON(http.StatusBadRequest, models.ErrBadRequest)
		return
	}
	err = uc.us.Insert(&json)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err":err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "user created successfully"})
}
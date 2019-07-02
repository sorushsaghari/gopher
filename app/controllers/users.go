package controllers

import (
	"../../app"
	"../../config"
	"../models"
	"../services"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type UserController struct {
	*app.App
}

func NewUserController(app *app.App) *UserController{

	return  &UserController{
		App: app,
	}
}

func (uc* UserController) SignIn(c* gin.Context)  {

	var cred config.Credentials
	_ = c.ShouldBindJSON(&cred)
	authenticated, err:= uc.User.Authenticate(cred.Email, cred.Password)
	uc.User.Find()
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
	me, err := uc.User.FindByEmail(claim.Email)
	if err != nil {
		c.JSON(404, err)
		return
	}
	c.JSON(200, me)
}



func (uc* UserController) Create(c* gin.Context) {
	var json services.UserDto
	c.ShouldBindJSON(&json)
	result, err := uc.App.User.Validate(json)
	if err != nil {
		fmt.Println(err)
		 c.JSON(http.StatusBadRequest, gin.H{"err":err.Error()})
		 return
	}
	if !result {
		c.JSON(http.StatusBadRequest, models.ErrBadRequest)
		return
	}
	err = uc.User.Insert(&json)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err":err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "user created successfully"})
}

func (uc* UserController) All(c* gin.Context) {
	users, err := uc.User.All()
	if err != nil{
		c.JSON(http.StatusBadRequest, models.ErrBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}


package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"github.com/tkanos/gonfig"
	"./models"
	"./config"
	"./middleware"
	"./controllers"
)

func main()  {
	configuration := config.Config{}
	err := gonfig.GetConf("./config/development-config.json", &configuration)
	services, err := models.NewService(
		models.WithGorm(configuration.Dialect, configuration.GetConnectionInfo()),
		models.WithUser(),
	)
	uc := controllers.NewUserController(services.User)
	if err != nil {
		panic(err)
	}
	defer services.Close()
	r := gin.Default()
	r.POST("/auth", uc.SignIn)
	r.POST("/", uc.Create)
	r.Use(middleware.IsAuthenticated())
	r.GET("/me", uc.Me)
	s := &http.Server{
		Addr:           ":3000",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
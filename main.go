package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"github.com/tkanos/gonfig"
	"./config"
	"./middleware"
	"./controllers"
	"./services"
)

func main()  {
	configuration := config.Config{}
	err := gonfig.GetConf("./config/development-config.json", &configuration)
	service, err := services.NewService(
		services.WithGorm(configuration.Dialect, configuration.GetConnectionInfo()),
		services.WithUser(),
	)
	uc := controllers.NewUserController(service.User)
	if err != nil {
		panic(err)
	}
	defer service.Close()
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
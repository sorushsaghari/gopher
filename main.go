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
		services.WithChat(),
	)
	uc := controllers.NewUserController(service.User)
	cc := controllers.NewChatController(service.Chat)
	if err != nil {
		panic(err)
	}
	defer service.Close()
	fs := http.FileServer(http.Dir("../public"))
	http.Handle("/", fs)
	r := gin.Default()
	private := r.Group("users")
	r.POST("/auth", uc.SignIn)
	r.POST("/users", uc.Create)
	r.GET("/users", uc.All)
	r.GET("/ws", cc.HandleConnection)
	private.Use(middleware.IsAuthenticated())
	private.GET("/me", uc.Me)
	s := &http.Server{
		Addr:           ":8000",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
	go cc.HandleMessages()
}
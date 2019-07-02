package main

import (
	"./app"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"github.com/tkanos/gonfig"
	"./config"
	"./app/middleware"
	"./app/controllers"
	"./app/services"
)

func main()  {

	configuration := config.Config{}
	err := gonfig.GetConf("./config/development-config.json", &configuration)
	service, err := services.WithGorm(configuration.Dialect, configuration.GetConnectionInfo())

		app, err := app.NewApp(
		app.WithUser(*service),
		app.WithChat(*service),
	)
	uc := controllers.NewUserController(app)
	cc := controllers.NewChatController(app)
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
	go cc.HandleMessages()

	s.ListenAndServe()
}
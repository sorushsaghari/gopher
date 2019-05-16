package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"github.com/tkanos/gonfig"
	"./models"
	"./config"
)

func main()  {
	configuration := config.Config{}
	err := gonfig.GetConf("./config/development-config.json", &configuration)
	services, err := models.NewService(
		models.WithGorm(configuration.Dialect, configuration.GetConnectionInfo()),
	)

	if err != nil {
		panic(err)
	}
	defer services.Close()
	 r := gin.Default()
	s := &http.Server{
		Addr:           ":3000",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
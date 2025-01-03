package main

import (
	"abramed_go/controller"
	"abramed_go/db"
	_ "abramed_go/docs"
	"abramed_go/repository"
	"abramed_go/routes"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type User struct {
	Nome     string `json:"nome" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func main() {
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1", "victor.controllab.com"})
	database, err := db.Init(".env")
	if err != nil {
		panic(err)
	}

	repository := repository.New(database)
	controller := controller.New(repository)

	routes.GenerateRouter(controller, r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	if err = r.Run("victor.controllab.com:8000"); err != nil {
		panic(err)
	}
}

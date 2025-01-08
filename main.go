package main

import (
	"abramed_go/controller"
	"abramed_go/db"
	_ "abramed_go/docs"
	"abramed_go/middlewares"
	"abramed_go/repository"
	"abramed_go/routes"
	"flag"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type User struct {
	Nome     string `json:"nome" binding:"required"`
	Password string `json:"password" binding:"required"`
}

var (
	PORT int
	HOST string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	MODE := os.Getenv("GIN_MODE")
	if MODE == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	
	flag.StringVar(&HOST, "host", "victor.controllab.com", "Porta para rodar o servidor")
	flag.IntVar(&PORT, "port", 8000, "Porta para rodar o servidor")
	flag.Parse()
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
	
	r.Use(middlewares.CorsMiddleware)
	routes.GenerateRouter(controller, r)
	
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	
	if err = r.Run(fmt.Sprintf("%s:%d", HOST, PORT)); err != nil {
		panic(err)
	}
}

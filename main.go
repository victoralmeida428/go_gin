package main

import (
	"abramed_go/controller"
	"abramed_go/db"
	"abramed_go/helpers"
	"abramed_go/middlewares"
	"abramed_go/repository"
	"abramed_go/routes"
	"flag"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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
	} else if MODE == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	
	flag.StringVar(&HOST, "host", "victor.controllab.com", "Porta para rodar o servidor")
	flag.IntVar(&PORT, "port", 8000, "Porta para rodar o servidor")
	flag.Parse()

	if err := helpers.UpdateJson(fmt.Sprintf("http://%s:%d/api/v1", HOST, PORT)); err != nil {
		panic(err)
	}
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

	//static files
	r.Static("/swagger-ui","./docs/swagger-ui")
	r.StaticFile("/swagger/doc.json","./docs/swagger.json")

	
	r.GET("/swagger", func (c *gin.Context){
		c.File("./docs/index.html")
	})
	
	r.Use(middlewares.CorsMiddleware)
	routes.GenerateRouter(controller, r)
	
	
	if err = r.Run(fmt.Sprintf("%s:%d", HOST, PORT)); err != nil {
		panic(err)
	}
}

package main

import (
	"abramed_go/cmd/api/controller"
	"abramed_go/cmd/api/db"
	"abramed_go/cmd/api/helpers"
	"abramed_go/cmd/api/middlewares"
	"abramed_go/cmd/api/repository"
	"abramed_go/cmd/api/routes"
	"flag"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


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
	r.Static("/swagger-ui","./cmd/api/docs/swagger-ui")
	r.StaticFile("/swagger/doc.json","./cmd/api/docs/swagger.json")

	
	r.GET("/swagger", func (c *gin.Context){
		c.File("./cmd/api/docs/index.html")
	})
	
	r.Use(middlewares.CorsMiddleware)

	routes.GenerateRouter(controller, r)
	
	
	if err = r.Run(fmt.Sprintf("%s:%d", HOST, PORT)); err != nil {
		panic(err)
	}
}

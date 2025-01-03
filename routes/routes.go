package routes

import (
	"abramed_go/controller"
	"abramed_go/middlewares"
	"github.com/gin-gonic/gin"
)

type Route struct {
	Path    string
	Method  string
	Handler func(ctx *gin.Context)
	Auth    bool
}

func GenerateRouter(c *controller.Controller, r *gin.Engine) {
	paths := userRoutes(c)
	paths = append(paths, formRoutes(c)...)
	paths = append(paths, variaveisRoutes(c)...)
	root := "/api/v1"
	for _, path := range paths {
		if path.Auth {
			r.Handle(path.Method, root+path.Path, middlewares.AuthenticationMiddleware, path.Handler)
		} else {
			r.Handle(path.Method, root+path.Path, path.Handler)
		}

	}

}

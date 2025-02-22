package routes

import (
	"abramed_go/cmd/api/controller"
	"abramed_go/cmd/api/middlewares"

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
	paths = append(paths, agendamentoRoutes(c)...)
	paths = append(paths, respsotasRoutes(c)...)
	root := "/api/v1"
	for _, path := range paths {
		if path.Auth {
			r.Handle(path.Method, root+path.Path, middlewares.AuthenticationMiddleware, path.Handler)
		} else {
			r.Handle(path.Method, root+path.Path, middlewares.LogAcesso, path.Handler)
		}

	}

}

package routes

import (
	"abramed_go/cmd/api/controller"
	"net/http"
)

func variaveisRoutes(ctl *controller.Controller) []Route {
	return []Route{
		{
			Path:    "/variable/type",
			Method:  http.MethodGet,
			Handler: ctl.Variavel.ListTipos,
			Auth:    true,
		},
		{
			Path:    "/variable/groups",
			Method:  http.MethodGet,
			Handler: ctl.Variavel.ListGrupo,
			Auth:    true,
		},
		{
			Path:    "/variable/groups",
			Method:  http.MethodPut,
			Handler: ctl.Variavel.CreateGrupo,
			Auth:    true,
		},
		{
			Path:    "/variable",
			Method:  http.MethodPut,
			Handler: ctl.Variavel.CreateVariavel,
			Auth:    true,
		},
		{
			Path:    "/variable",
			Method:  http.MethodGet,
			Handler: ctl.Variavel.ListVariavel,
			Auth:    true,
		},
	}
}

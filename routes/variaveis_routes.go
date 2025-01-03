package routes

import (
	"abramed_go/controller"
	"net/http"
)

func variaveisRoutes(ctl *controller.Controller) []Route {
	return []Route{
		{
			Path:    "/variaveis/tipo",
			Method:  http.MethodGet,
			Handler: ctl.Variavel.ListTipos,
			Auth:    true,
		},
		{
			Path:    "/variaveis/grupamento",
			Method:  http.MethodGet,
			Handler: ctl.Variavel.ListGrupo,
			Auth:    true,
		},
		{
			Path:    "/variaveis/grupamento",
			Method:  http.MethodPut,
			Handler: ctl.Variavel.CreateGrupo,
			Auth:    true,
		},
		{
			Path:    "/variaveis",
			Method:  http.MethodPut,
			Handler: ctl.Variavel.CreateVariavel,
			Auth:    true,
		},
		{
			Path:    "/variaveis",
			Method:  http.MethodGet,
			Handler: ctl.Variavel.ListVariavel,
			Auth:    true,
		},
		{
			Path:    "/grupamento",
			Method:  http.MethodGet,
			Handler: ctl.Variavel.ListGrupamentos,
			Auth:    true,
		},
	}
}

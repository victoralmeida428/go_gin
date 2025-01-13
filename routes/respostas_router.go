package routes

import (
	"abramed_go/controller"
	"net/http"
)

func respsotasRoutes(ctl *controller.Controller) []Route {
	return []Route{
		{
			Path:    "/answer",
			Method:  http.MethodPost,
			Handler: ctl.Resposta.Responder,
			Auth:    true,
		},
		{
			Path:    "/answer",
			Method:  http.MethodGet,
			Handler: ctl.Resposta.ListRepostas,
			Auth:    true,
		},
	}
}

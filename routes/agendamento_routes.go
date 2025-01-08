package routes

import (
	"abramed_go/controller"
	"net/http"
)

func agendamentoRoutes(ctl *controller.Controller) []Route {
	return []Route{
		{
			Path:    "/schedule",
			Method:  http.MethodGet,
			Handler: ctl.Agendamento.ListAll,
			Auth:    true,
		},
	}
}

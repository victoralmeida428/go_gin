package routes

import (
	"abramed_go/cmd/api/controller"
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
		{
			Path:    "/schedule",
			Method:  http.MethodPut,
			Handler: ctl.Agendamento.Create,
			Auth:    true,
		},
		{
			Path:    "/schedule/periodicity",
			Method:  http.MethodGet,
			Handler: ctl.Agendamento.ListPeriodicity,
			Auth:    true,
		},
	}
}

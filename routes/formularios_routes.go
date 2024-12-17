package routes

import (
	"abramed_go/controller"
	"net/http"
)

func formRoutes(ctl *controller.Controller) []Route {
	return []Route{
		{
			Path:    "/formulario",
			Method:  http.MethodGet,
			Handler: ctl.Form.List,
			Auth:    true,
		},
		{
			Path:    "/formulario",
			Method:  http.MethodPut,
			Handler: ctl.Form.Create,
			Auth:    true,
		},
	}
}

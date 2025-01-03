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
		{
			Path: "/formulario/:id",
			Method: http.MethodDelete,
			Handler: ctl.Form.Delete,
			Auth: true,
		},
		{
			Path: "/formulario",
			Method: http.MethodPatch,
			Handler: ctl.Form.Update,
			Auth: true,
		},
	}
}

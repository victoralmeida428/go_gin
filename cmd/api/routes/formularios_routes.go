package routes

import (
	"abramed_go/cmd/api/controller"
	"net/http"
)

func formRoutes(ctl *controller.Controller) []Route {
	return []Route{
		{
			Path:    "/form",
			Method:  http.MethodGet,
			Handler: ctl.Form.List,
			Auth:    true,
		},
		{
			Path:    "/form",
			Method:  http.MethodPut,
			Handler: ctl.Form.Create,
			Auth:    true,
		},
		{
			Path:    "/form/:id",
			Method:  http.MethodDelete,
			Handler: ctl.Form.Delete,
			Auth:    true,
		},
		{
			Path:    "/form/:id",
			Method:  http.MethodGet,
			Handler: ctl.Form.FindById,
			Auth:    true,
		},
		{
			Path:    "/form",
			Method:  http.MethodPatch,
			Handler: ctl.Form.Update,
			Auth:    true,
		},
		{
			Path:    "/form/methods",
			Method:  http.MethodGet,
			Handler: ctl.Form.ListMethods,
			Auth:    true,
		},
	}
}

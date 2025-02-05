package routes

import (
	"abramed_go/cmd/api/controller"
	"net/http"
)

func userRoutes(ctl *controller.Controller) []Route {
	return []Route{
		{
			Path:    "/user/login",
			Method:  http.MethodPost,
			Handler: ctl.User.Login,
		},
		{
			Path:    "/user",
			Method:  http.MethodGet,
			Handler: ctl.User.GetUser,
			Auth:    true,
		},
		{
			Path:    "/user/create",
			Method:  http.MethodPut,
			Handler: ctl.User.CreateUser,
			Auth:    false,
		},
		{
			Path:    "/user/groups",
			Method:  http.MethodPost,
			Handler: ctl.User.AddGrupamento,
			Auth:    true,
		},
		{
			Path:    "/user/groups",
			Method:  http.MethodDelete,
			Handler: ctl.User.DelGrupamento,
			Auth:    true,
		},
	}
}

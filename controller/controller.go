package controller

import (
	"abramed_go/repository"
)

type Controller struct {
	User *userController
	Form *formController
	Variavel *variavelController
}

func New(repo *repository.Repository) *Controller {
	return &Controller{
		User: &userController{repo: repo},
		Form: &formController{repo: repo},
		Variavel: &variavelController{repo: repo},
	}
}

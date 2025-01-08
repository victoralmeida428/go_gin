package controller

import (
	"abramed_go/repository"
)

type Controller struct {
	User *userController
	Form *formController
	Variavel *variavelController
	Agendamento *agendamentoController
}

func New(repo *repository.Repository) *Controller {
	return &Controller{
		User: &userController{repo: repo},
		Form: &formController{repo: repo},
		Variavel: &variavelController{repo: repo},
		Agendamento: &agendamentoController{repo: repo},
	}
}

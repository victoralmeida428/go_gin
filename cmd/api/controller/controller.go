package controller

import (
	"abramed_go/cmd/api/repository"
)

type Controller struct {
	User        *userController
	Form        *formController
	Variavel    *variavelController
	Agendamento *agendamentoController
	Resposta    *respostaController
}

func New(repo *repository.Repository) *Controller {
	return &Controller{
		User:        &userController{repo: repo},
		Form:        &formController{repo: repo},
		Variavel:    &variavelController{repo: repo},
		Agendamento: &agendamentoController{repo: repo},
		Resposta:    &respostaController{repo: repo},
	}
}

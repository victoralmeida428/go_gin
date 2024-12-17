package controller

import (
	"abramed_go/repository"
)

type Controller struct {
	User *userController
	Form *formController
}

func New(repo *repository.Repository) *Controller {
	return &Controller{
		User: &userController{userUsecase: repo.User},
		Form: &formController{formRepo: repo.Formulario},
	}
}

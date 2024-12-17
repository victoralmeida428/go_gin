package repository

import (
	"abramed_go/model"
	"database/sql"
)

type Repository struct {
	User       IUserRepository
	Formulario IRepository[model.Formulario]
}

func New(db *sql.DB) *Repository {
	return &Repository{
		User:       &UserRepository{connection: db},
		Formulario: &FormularioRepository{connection: db},
	}

}

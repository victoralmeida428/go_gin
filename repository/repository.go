package repository

import (
	"abramed_go/model"
	"database/sql"
)

type Repository struct {
	User       interface {
		IRepository[model.User]
		FindByUsuarioSenha(usuario, senha string) (model.User, error)
		AddGrupamento(usuario *model.User, grupo *model.Grupamento) error
		FindAllMembers(id int) ([]model.User, error)
	}
	Formulario IRepository[model.Formulario]
	Variavel interface {
		IRepository[model.Variavel]
		FindGrupamentoById(id int) (*model.Grupamento, error)
		CreateGrupamento(grupo *model.Grupamento) (*model.Grupamento, error)
		ListTipos() ([]model.TipoVariavel, error)
		ListGrupamentos() ([]model.Grupamento, error)
		ListGrupamentosByUser(user *model.User) ([]model.Grupamento, error)
	}
}

func New(db *sql.DB) *Repository {
	return &Repository{
		User:       &UserRepository{connection: db},
		Formulario: &FormularioRepository{connection: db},
		Variavel: &VariavelRepository{connection: db},
		
	}

}

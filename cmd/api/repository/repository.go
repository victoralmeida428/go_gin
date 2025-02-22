package repository

import (
	"database/sql"
)

type Repository struct {
	User       *UserRepository
	Formulario *FormularioRepository
	Variavel   *VariavelRepository
	Agendamento *AgendamentoRepository
	Resposta *RespostaRepository
}

func New(db *sql.DB) *Repository {
	return &Repository{
		User:       &UserRepository{connection: db},
		Formulario: &FormularioRepository{connection: db},
		Variavel:   &VariavelRepository{connection: db},
		Agendamento: &AgendamentoRepository{connection: db},
		Resposta: &RespostaRepository{connection: db},
		
	}

}

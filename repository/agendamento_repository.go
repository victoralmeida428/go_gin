package repository

import (
	"abramed_go/dto/dto_agendamento"
	"abramed_go/model"
	"database/sql"
)

type AgendamentoRepository struct {
	connection *sql.DB
}

func (repo *AgendamentoRepository) FindAll(user *model.User) ([]model.Agendamento, error) {
	query := `
	SELECT
		a.id,
		a.inicio,
		a.inicio + p.intervalo proximo,
		a.ativo,
		metodo_id,
		m.nome metodo,
		a.periodicidade_id, 
		p.nome periodicidade,
		f.id id_formulario,
		f.nome formulario,
		vf.versao versao_formulario,
		vf.id versao_formulario_id,
		us.id user_id,
		us.first_name,
		us.last_name,
		us.email,
		f.descricao
	FROM
		indicadores.agendamento a
		JOIN indicadores.periodicidade p ON p.id = a.periodicidade_id
		JOIN indicadores.versao_formulario vf ON vf.id = a.versao_formulario_id
		JOIN indicadores.formulario f ON f.id = vf.formulario_id
		JOIN indicadores.usuario us ON us.id = a.user_id
		JOIN indicadores.metodo m ON m.id = a.metodo_id
	WHERE
		us.empresa_id = $1
	`
	result, err := repo.connection.Query(query, user.EmpresaID.ID)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	agendamentos := make([]model.Agendamento, 0)
	for result.Next() {
		var agendamento model.Agendamento
		err = result.Scan(
			&agendamento.ID,
			&agendamento.Inicio,
			&agendamento.Proximo,
			&agendamento.Ativo,
			&agendamento.Metodo.ID,
			&agendamento.Metodo.Nome,
			&agendamento.Periodicidade.ID,
			&agendamento.Periodicidade.Nome,
			&agendamento.VersaoFormulario.Formulario.ID,
			&agendamento.VersaoFormulario.Formulario.Nome,
			&agendamento.VersaoFormulario.Versao,
			&agendamento.VersaoFormulario.ID,
			&agendamento.User.ID,
			&agendamento.User.FirstName,
			&agendamento.User.LastName,
			&agendamento.User.Email,
			&agendamento.VersaoFormulario.Formulario.Descricao,
		)
		if err != nil {
			return nil, err
		}
		agendamentos = append(agendamentos, agendamento)
	}
	return agendamentos, nil
}

func (repo *AgendamentoRepository) Insert(userId int, schedule dto_agendamento.CreateInput) error {
	query := `
		INSERT INTO indicadores.agendamento (versao_formulario_id, user_id, periodicidade_id, inicio, ativo, metodo_id)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	tx, err := repo.connection.Begin()
	if err != nil {
		return err
	}
	args := []interface{}{schedule.Versao, userId, schedule.Periodicidade, schedule.Inicio, schedule.Ativo, schedule.Metodo}
	_, err = tx.Exec(query, args...)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (repo *AgendamentoRepository) ListPeriodicity() ([]model.Periodicidade, error) {
	query := `
		SELECT
			id, nome
		FROM
			indicadores.periodicidade
		ORDER BY intervalo
	`
	results, err := repo.connection.Query(query)
	if err != nil {
		return nil, err
	}

	periodos := make([]model.Periodicidade, 0)
	for results.Next() {
		var periodo model.Periodicidade
		err := results.Scan(
			&periodo.ID,
			&periodo.Nome,
		)
		if err != nil {
			return nil, err
		}
		periodos = append(periodos, periodo)
	}
	return periodos, nil
}

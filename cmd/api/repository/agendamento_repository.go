package repository

import (
	"abramed_go/cmd/api/dto/dto_agendamento"
	"abramed_go/cmd/api/helpers"
	"abramed_go/cmd/api/helpers/status"
	"abramed_go/cmd/api/model"
	"database/sql"
	"time"
)

type AgendamentoRepository struct {
	connection *sql.DB
}

func (repo *AgendamentoRepository) FindAll(user *model.User, statusForm *status.StatusForm) ([]model.Agendamento, error) {
	query := `
	SELECT
		-- *
		a.id,
		a.inicio,
		a.inicio + p.intervalo + '-1 day' "end",
		NOW() BETWEEN inicio + intervalo AND inicio  + 2 * intervalo + '-1 day' ativo,
		MAX(metodo_id) metodo_id,
		MAX(m.nome) metodo,
		MAX(a.periodicidade_id),
		MAX(p.nome) periodicidade,
		MAX(f.id) id_formulario,
		MAX(f.nome) formulario,
		MAX(vf.versao) versao_formulario,
		MAX(vf.id) versao_formulario_id,
		MAX(us.id) usuario_id,
		MAX(us.nome) nome,
		MAX(us.email) email,
		MAX(f.descricao) descricao,
		a.inicio + 2 * p.intervalo + '-1 day' prazo,
		-- Data da resposta mais recente das perguntas optativas
        MAX(resposta.criado_em) FILTER ( WHERE not variavel.obrigatorio ) resposta_optativas_mais_recente,
		-- Data da resposta mais recente das perguntas obrigatórias
        MAX(resposta.criado_em) FILTER ( WHERE variavel.obrigatorio ) resposta_obrigatoria_mais_recente,
        -- Optativa total
        COUNT(1) FILTER ( WHERE NOT variavel.obrigatorio ) opt_total,
        -- Não respondido optativo
        COUNT(1) FILTER ( WHERE resposta.id NOTNULL AND NOT variavel.obrigatorio )  opt,
        -- Total obrigatório
        COUNT(1) FILTER ( WHERE variavel.obrigatorio ) obrigatorio_total,
        -- Não respondido obrigatório
        COUNT(1) FILTER ( WHERE resposta.id NOTNULL AND variavel.obrigatorio )  obrigatorio
	FROM
		indicadores.agendamento a
		JOIN indicadores.periodicidade p ON p.id = a.periodicidade_id
		JOIN indicadores.usuario us ON us.id = a.usuario_id
		JOIN indicadores.versao_formulario vf ON vf.id = a.versao_formulario_id
		JOIN indicadores.formulario f ON f.id = vf.formulario_id
		JOIN indicadores.metodo m ON m.id = a.metodo_id
		JOIN indicadores.variavel ON vf.id = a.versao_formulario_id
		LEFT JOIN indicadores.resposta ON a.id = resposta.agendamento_id
		AND resposta.variavel_id = variavel.id
	WHERE
		us.id = $1
		AND a.inicio + p.intervalo <= NOW()
		AND NOT variavel.possui_item
	GROUP BY
		a.id,
		a.usuario_id,
		p.intervalo,
		a.inicio
	ORDER BY
		a.inicio DESC;
	`
	result, err := repo.connection.Query(query, user.ID)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	agendamentos := make([]model.Agendamento, 0)
	for result.Next() {
		var required, opt, requiredTotal, optTotal int
		var requiredTime, optTime *time.Time
		var agendamento model.Agendamento
		err = result.Scan(
			&agendamento.ID,
			&agendamento.Inicio,
			&agendamento.End,
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
			&agendamento.User.Name,
			&agendamento.User.Email,
			&agendamento.VersaoFormulario.Formulario.Descricao,
			&agendamento.Deadline,
			&optTime,
			&requiredTime,
			&optTotal,
			&opt,
			&requiredTotal,
			&required,
		)
		dataDeafult := time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local)
		if optTime == nil {
			optTime = &dataDeafult
		}
		if requiredTime == nil {
			requiredTime = &dataDeafult
		}

		stat := status.GetStatusForm(requiredTotal, optTotal, required, opt, agendamento.Deadline, *optTime, *requiredTime)
		if statusForm != nil && stat != *statusForm {
			continue
		}
		agendamento.Status = stat

		if err != nil {
			return nil, err
		}
		year, week := helpers.GetWeek(agendamento.Inicio)
		agendamento.Week = &week
		agendamento.Year = &year
		agendamentos = append(agendamentos, agendamento)
	}
	return agendamentos, nil
}

func (repo *AgendamentoRepository) Insert(userId int, schedule *dto_agendamento.CreateInput) error {
	query := `
		INSERT INTO indicadores.agendamento (versao_formulario_id, usuario_id, periodicidade_id, inicio, ativo, metodo_id)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	tx, err := repo.connection.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	args := []interface{}{schedule.Versao, userId, schedule.Periodicidade, schedule.Inicio, schedule.Ativo, schedule.Metodo}
	_, err = tx.Exec(query, args...)
	tx.Commit()
	return err
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

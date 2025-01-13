package repository

import (
	"abramed_go/dto/dto_respostas"
	"abramed_go/model"
	"database/sql"
	"fmt"
)

type RespostaRepository struct {
	connection *sql.DB
}

func (repo *RespostaRepository) GetVersion(resposta int) (versao int, err error) {
	query := `
	SELECT
		count(id)+1
	FROM
		indicadores.versao_resposta_numerica
	WHERE
		resposta_id = $1
	`

	err = repo.connection.QueryRow(query, resposta).Scan(&versao)
	return
}

func (repo *RespostaRepository) VerificarRespondido(variavel, user, agenda int) (int, error) {
	query := `
	SELECT
		id
	FROM
		indicadores.resposta
	WHERE
		variavel_id=$1 and usuario_id=$2 and agendamento_id=$3
	`
	var contagem int
	if err := repo.connection.QueryRow(query, variavel, user, agenda).Scan(&contagem); err != nil {
		return 0, err
	}
	return contagem, nil
}

func (repo *RespostaRepository) Responder(input dto_respostas.ResponderInput, userId int) error {
	tx, err := repo.connection.Begin()
	if err != nil {
		return err
	}
	respostaId, err := repo.VerificarRespondido(input.Variavel, userId, input.Agendamento)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(respostaId)
	if respostaId ==0 {
		query_resp := `
		INSERT INTO indicadores.resposta (variavel_id, agendamento_id, usuario_id)
		VALUES ($1, $2, $3)
		RETURNING id
		`
		if err = tx.QueryRow(query_resp, input.Variavel, input.Agendamento, userId).Scan(&respostaId); err != nil {
			tx.Rollback()
			return err
		}
	}
	fmt.Println(respostaId)
	query_valor := `
	INSERT INTO indicadores.resposta_numerica (valor, ativo)
	VALUES ($1, $2)
	RETURNING id
	`
	var valorId int

	if err = tx.QueryRow(query_valor, input.Valor, input.Ativo).Scan(&valorId); err != nil {
		tx.Rollback()
		return err
	}

	versao, err := repo.GetVersion(respostaId)
	if err != nil {
		tx.Rollback()
		return err
	}

	query_version := `
	INSERT INTO indicadores.versao_resposta_numerica (resposta_numerica_id, resposta_id, versao)
	VALUES ($1, $2, $3)
	`
	if _, err = tx.Exec(query_version, valorId, respostaId, versao); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (repo *RespostaRepository) ListAll(id int) ([]model.Resposta, error) {
	query := `
	SELECT
	variavel.id variavael_id,
	usuario.id,
	usuario.first_name,
	usuario.last_name,
	resposta_numerica.valor,
	versao_resposta_numerica.versao,
	versao_resposta_numerica.criado_em,
	resposta_numerica.ativo
FROM
	indicadores.variavel_formulario varf
	JOIN indicadores.view_variavel variavel ON varf.variavel_id = variavel.id
	JOIN indicadores.tipo_variavel tipo ON tipo.id = tipo_variavel_id
	JOIN indicadores.agendamento agenda ON agenda.versao_formulario_id = varf.versao_formulario_id
	JOIN indicadores.resposta ON resposta.agendamento_id = agenda.id and resposta.variavel_id = varf.variavel_id
	JOIN indicadores.versao_resposta_numerica ON versao_resposta_numerica.resposta_id = resposta.id
	JOIN indicadores.resposta_numerica ON resposta_numerica.id = versao_resposta_numerica.resposta_numerica_id
	JOIN indicadores.usuario ON usuario.id = resposta.usuario_id
WHERE
	agenda.id = $1
	`
	results, err := repo.connection.Query(query, id)
	if err != nil {
		return nil, err
	}
	respostas := make([]model.Resposta, 0)
	indexMap := make(map[int]int)
	count := 0

	for results.Next() {
		var resposta model.Resposta
		var respostasNum model.RespostaNum
		if err := results.Scan(
			&resposta.Variavel.ID,
			&resposta.Usuario.ID,
			&resposta.Usuario.FirstName,
			&resposta.Usuario.LastName,
			&respostasNum.Valor,
			&respostasNum.Versao,
			&respostasNum.CriadoEm,
			&respostasNum.Ativo,
		); err != nil {
			return nil, err
		}


		if idx, exists := indexMap[resposta.Variavel.ID]; !exists {
			resposta.Respostas = []model.RespostaNum{respostasNum}
			respostas = append(respostas, resposta)
		} else {
			respostas[idx].Respostas = append(respostas[idx].Respostas, respostasNum) 
		}
		indexMap[resposta.Variavel.ID] = count
		count++

	}

	return respostas, nil
}
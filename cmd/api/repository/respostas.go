package repository

import (
	"abramed_go/cmd/api/dto/dto_respostas"
	"abramed_go/cmd/api/model"
	"database/sql"
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

	args := []interface{}{input.Agendamento, input.Variavel, userId, input.Valor}

	query_valor := `
	INSERT INTO
	indicadores.resposta (
		agendamento_id,
		variavel_id,
		usuario_id,
		resposta_numerica
	)
	VALUES
		($1, $2, $3, $4)
	ON CONFLICT (agendamento_id, variavel_id) DO
	UPDATE
	SET
		resposta_numerica = excluded.resposta_numerica,
		usuario_id = excluded.usuario_id,
		criado_em = now()
	`

	if _, err = tx.Exec(query_valor, args...); err != nil {
		tx.Rollback()
		return err
	}

	query_log := `
	INSERT INTO
		indicadores.log_resposta_numerica (
			agendamento_id,
			variavel_id,
			usuario_id,
			resposta_numerica
		)
	VALUES
		($1, $2, $3, $4)
	`
	if _, err = tx.Exec(query_log, args...); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func (repo *RespostaRepository) ListAll(id int) ([]model.Resposta, error) {
	query := `
	SELECT
	variavel_id,
	usuario.id,
	usuario.nome,
	log_.resposta_numerica valor,
	log_.criado_em
FROM
	indicadores.log_resposta_numerica log_
	JOIN indicadores.usuario ON usuario.id = log_.usuario_id
WHERE
	agendamento_id = $1
ORDER BY
	variavel_id, criado_em DESC
	`
	results, err := repo.connection.Query(query, id)
	if err != nil {
		return nil, err
	}
	respostas := make([]model.Resposta, 0)
	indexMap := make(map[int]int)

	for results.Next() {
		var resposta model.Resposta
		var respostasNum model.RespostaNum
		if err := results.Scan(
			&resposta.Variavel.ID,
			&resposta.Usuario.ID,
			&resposta.Usuario.Name,
			&respostasNum.Valor,
			&respostasNum.CriadoEm,
		); err != nil {
			return nil, err
		}

		if idx, exists := indexMap[resposta.Variavel.ID]; !exists {
			resposta.Respostas = []model.RespostaNum{respostasNum}
			respostas = append(respostas, resposta)
			indexMap[resposta.Variavel.ID] = len(respostas) - 1
		} else {
			respostas[idx].Respostas = append(respostas[idx].Respostas, respostasNum)
		}
	}
	return respostas, nil
}

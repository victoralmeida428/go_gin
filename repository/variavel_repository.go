package repository

import (
	"abramed_go/model"
	"database/sql"
)

type VariavelRepository struct {
	connection *sql.DB
}

func (repo *VariavelRepository) FindAll() ([]model.ViewVariavel, error) {
	query := `
		SELECT variavel.id, tipo_variavel_id, tipo.nome, pergunta, item, possui_item, obrigatorio
		FROM indicadores.view_variavel variavel
		JOIN indicadores.tipo_variavel tipo ON tipo.id = tipo_variavel_id
	`
	result, err := repo.connection.Query(query)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	variaveis := make([]model.ViewVariavel, 0)
	for result.Next() {
		var variavel model.ViewVariavel
		err = result.Scan(&variavel.ID, &variavel.TipoVariavel.ID, &variavel.TipoVariavel.Nome, &variavel.Pergunta, &variavel.Item, &variavel.PossuiItem, &variavel.Obrigatorio)
		if err != nil {
			return nil, err
		}
		variaveis = append(variaveis, variavel)
	}
	return variaveis, nil
}

func (form *VariavelRepository) FindById(id int) (model.Variavel, error) {
	var variavel model.Variavel
	return variavel, nil
}

func (form *VariavelRepository) Update(variavel *model.Variavel) (*model.Variavel, error) {
	return nil, nil
}

func (form *VariavelRepository) Delete(id int) error {
	return nil
}

func (form *VariavelRepository) Insert(variavel *model.Variavel) error {
	query := `
	INSERT INTO indicadores.variavel(grupamento_id, tipo_variavel_id, pergunta_id, possui_item, obrigatorio, texto)
	VALUES ($1, $2, $3, $4, $5, $6)
	`

	args := []interface{}{ variavel.TipoVariavel, variavel.PerguntaId, variavel.PossuiItem, variavel.Obrigatorio, variavel.Texto}
	_, err := form.connection.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (form *VariavelRepository) CreateGrupamento(grupo *model.Grupamento) (*model.Grupamento, error) {
	tx, err := form.connection.Begin()
	if err != nil {
		return nil, err
	}

	query := `
	INSERT INTO indicadores.grupamento(nome)
	VALUES ($1)
	RETURNING id
	`

	args := []interface{}{grupo.Nome}
	var Id int
	err = tx.QueryRow(query, args...).Scan(&Id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return form.FindGrupamentoById(Id)

}

func (form *VariavelRepository) FindGrupamentoById(id int) (*model.Grupamento, error) {
	query := `
	SELECT id, nome FROM indicadores.grupamento WHERE id=$1
	`

	var grupo model.Grupamento
	if err := form.connection.QueryRow(query, id).Scan(&grupo.ID, &grupo.Nome); err != nil {
		return nil, err
	}

	return &grupo, nil
}

func (form *VariavelRepository) ListTipos() ([]model.TipoVariavel, error) {
	query := `
	SELECT id, nome FROM indicadores.tipo_variavel
	`
	result, err := form.connection.Query(query)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	tipos := make([]model.TipoVariavel, 0)
	for result.Next() {
		var tipo model.TipoVariavel
		err := result.Scan(&tipo.ID, &tipo.Nome)
		if err != nil {
			return nil, err
		}
		tipos = append(tipos, tipo)
	}

	return tipos, nil
}

func (form *VariavelRepository) ListGrupamentos() ([]model.Grupamento, error) {
	query := `
	SELECT id, nome FROM indicadores.grupamento
	`
	result, err := form.connection.Query(query)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	grupos := make([]model.Grupamento, 0)
	for result.Next() {
		var grupo model.Grupamento
		err := result.Scan(&grupo.ID, &grupo.Nome)
		if err != nil {
			return nil, err
		}
		grupos = append(grupos, grupo)
	}

	return grupos, nil
}

func (form *VariavelRepository) ListGrupamentosByUser(user *model.User) ([]model.Grupamento, error) {
	query := `
	SELECT gr.id, gr.nome FROM indicadores.grupamento gr
	JOIN indicadores.grupamento_usuario ugr ON gr.id = ugr.grupamento_id
	JOIN indicadores.usuario usr ON ugr.usuario_id = usr.id
	where usr.id=$1
	`
	result, err := form.connection.Query(query, user.ID)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	grupos := make([]model.Grupamento, 0)
	for result.Next() {
		var grupo model.Grupamento
		err := result.Scan(&grupo.ID, &grupo.Nome)
		if err != nil {
			return nil, err
		}
		grupos = append(grupos, grupo)
	}

	return grupos, nil
}

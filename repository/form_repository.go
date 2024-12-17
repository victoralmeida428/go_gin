package repository

import (
	"abramed_go/model"
	"database/sql"
)

type FormularioRepository struct {
	connection *sql.DB
}

func (form *FormularioRepository) FindAll() ([]model.Formulario, error) {
	query := `
	SELECT id, nome, descricao, ativo
	FROM indicadores.formulario
	`
	rows, err := form.connection.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var formularios []model.Formulario
	for rows.Next() {
		var formulario model.Formulario
		if err = rows.Scan(&formulario.ID, &formulario.Nome, &formulario.Descricao, &formulario.Ativo); err != nil {
			return nil, err
		}
		formularios = append(formularios, formulario)
	}
	return formularios, nil
}

func (form *FormularioRepository) FindById(id int) (model.Formulario, error) {
	query := `
	SELECT id, nome, descricao, ativo
	FROM indicadores.formulario
	where id = $1
	`
	var formulario model.Formulario
	rows, err := form.connection.Query(query, id)
	if err != nil {
		return formulario, err
	}
	for rows.Next() {
		if err = rows.Scan(&formulario.ID, &formulario.Nome, &formulario.Descricao, &formulario.Ativo); err != nil {
			return formulario, err
		}
	}
	return formulario, nil
}

func (form *FormularioRepository) Update(formulario *model.Formulario) (*model.Formulario, error) {
	return formulario, nil
}

func (form *FormularioRepository) Insert(formulario *model.Formulario) error {
	query := `
		INSERT INTO indicadores.formulario (nome, descricao, ativo)
		VALUES ($1, $2, $3)
		RETURNING id, nome, descricao, ativo;
		`
	args := []interface{}{formulario.Nome, formulario.Descricao, formulario.Ativo}
	rows, err := form.connection.Query(query, args...)
	if err != nil {
		return err
	}

	for rows.Next() {
		if err = rows.Scan(&formulario.ID, &formulario.Nome, &formulario.Descricao, &formulario.Ativo); err != nil {
			return err
		}
	}

	query_versao := `
		INSERT INTO indicadores.versao_formulario (formulario_id, versao)
		VALUES ($1, $2)
		`
	args = []interface{}{formulario.ID, int(1)}
	_, err = form.connection.Exec(query_versao, args...)
	return err
}

func (form *FormularioRepository) Delete(formulario *model.Formulario) error { return nil }

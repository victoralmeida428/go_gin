package repository

import (
	"abramed_go/model"
	"database/sql"
	"errors"
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
	query := `
	UPDATE indicadores.formulario
	SET nome = $1,
		descricao = $2,
		ativo = $3
	where id = $4
	`
	args := []interface{}{formulario.Nome, formulario.Descricao, formulario.Ativo, formulario.ID}
	results, err := form.connection.Exec(query, args...)
	if err != nil {
		return formulario, err
	}
	rows, err := results.RowsAffected()
	if err != nil {
		return formulario, err
	}
	if rows==0{
		return formulario, errors.New("id not found")
	}

	return formulario, nil
}

func (form *FormularioRepository) Insert(formulario *model.Formulario) error {
	tx, err := form.connection.Begin()
	if err != nil {
		return err
	}
	query := `
		INSERT INTO indicadores.formulario (nome, descricao, ativo)
		VALUES ($1, $2, $3)
		RETURNING id, nome, descricao, ativo;
		`
	args := []interface{}{formulario.Nome, formulario.Descricao, formulario.Ativo}
	rows, err := tx.Query(query, args...)
	if err != nil {
		tx.Rollback()
		return err
	}

	for rows.Next() {
		if err = rows.Scan(&formulario.ID, &formulario.Nome, &formulario.Descricao, &formulario.Ativo); err != nil {
			tx.Rollback()
			return err
		}
	}

	query_versao := `
		INSERT INTO indicadores.versao_formulario (formulario_id, versao)
		VALUES ($1, $2)
		`
	args = []interface{}{formulario.ID, int(1)}
	_, err = tx.Exec(query_versao, args...)
	if err != nil {
		tx.Rollback()
	}

	tx.Commit()
	return nil
}

func (form *FormularioRepository) Delete(id int) error { 
	tx, err := form.connection.Begin()
	if err != nil {
		return err
	}
	query := `
		DELETE FROM indicadores.versao_formulario where formulario_id=$1
	`
	result, err := tx.Exec(query, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}
	if rows == 0 {
		tx.Rollback()
		return errors.New("formulario_id not found")
	}

	deleteForm := `
	DELETE FROM indicadores.formulario where id=$1
	
	`
	result, err = tx.Exec(deleteForm, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	rows, err = result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}
	if rows == 0 {
		tx.Rollback()
		return errors.New("id not found")
	}

	tx.Commit()
	return nil 
}

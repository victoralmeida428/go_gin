package repository

import (
	"abramed_go/helpers"
	"abramed_go/model"
	"database/sql"
	"errors"
)
type UserRepository struct {
	connection *sql.DB
}

func (ur *UserRepository) FindAll() ([]model.User, error) {
	query := `
	SELECT id, usuario, senha, email, manager_id
	FROM indicadores.usuario
	`
	rows, err := ur.connection.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []model.User
	for rows.Next() {
		var user model.User
		if err = rows.Scan(&user.ID, &user.Usuario, &user.Senha, &user.Email, &user.ManagerID); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (ur *UserRepository) FindAllMembers(id int) ([]model.User, error) {
	query := `
	SELECT id, usuario, senha, email, manager_id
	FROM indicadores.usuario WHERE manager_id = $1
	`
	rows, err := ur.connection.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []model.User
	for rows.Next() {
		var user model.User
		if err = rows.Scan(&user.ID, &user.Usuario, &user.Senha, &user.Email, &user.ManagerID); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (ur *UserRepository) FindById(id int) (model.User, error) {
	query := `
	SELECT id, usuario, senha, email, manager_id, empresa
	FROM indicadores.usuario
	where id = $1
	`
	var user model.User
	rows, err := ur.connection.Query(query, id)
	if err != nil {
		return user, err
	}
	for rows.Next() {
		if err = rows.Scan(&user.ID, &user.Usuario, &user.Senha, &user.Email, &user.ManagerID, &user.Empresa); err != nil {
			return user, err
		}

	}
	return user, nil
}

func (ur *UserRepository) Update(user *model.User) (*model.User, error) { return user, nil }

func (ur *UserRepository) Insert(user *model.User) error {
	query := `
		INSERT INTO indicadores.usuario(usuario, senha, email, empresa, manager_id)
		VALUES ($1, $2, $3, $4, $5)
		`
	hash, err := helpers.HashPassword(user.Senha)
	if err != nil {
		return err
	}
	args := []interface{}{user.Usuario, hash, user.Email, user.Empresa, user.ManagerID}
	_, err = ur.connection.Exec(query, args...)
	return err
}

func (ur *UserRepository) Delete(id int) error { return nil }

func (ur *UserRepository) FindByUsuarioSenha(usuario, senha string) (model.User, error) {
	query := `
	SELECT id, usuario, senha, email, manager_id, empresa
	FROM indicadores.usuario
	where usuario = $1
	`
	var user model.User
	rows, err := ur.connection.Query(query, usuario)
	if err != nil {
		return user, err
	}
	for rows.Next() {
		if err = rows.Scan(&user.ID, &user.Usuario, &user.Senha, &user.Email, &user.ManagerID, &user.Empresa); err != nil {
			return user, err
		}

	}
	if !helpers.CheckPasswordHash(senha, user.Senha) {
		return model.User{}, errors.New("senha inválida")
	}
	return user, nil

}

func (ur *UserRepository) AddGrupamento(usuario *model.User, grupo *model.Grupamento) error {
	tx, err := ur.connection.Begin()
	if err != nil {
		return err
	}
	query := `
	INSERT INTO indicadores.grupamento_usuario(usuario_id, grupamento_id)
	VALUES ($1, $2)
	`
	_, err = tx.Exec(query, usuario.ID, grupo.ID)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
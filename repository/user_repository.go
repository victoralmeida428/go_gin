package repository

import (
	"abramed_go/helpers"
	"abramed_go/model"
	"database/sql"
	"errors"
	"strings"
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
	SELECT usr.id, usuario, email, manager_id, empresa_id, empresa.nome, nivel.id, nivel.nome, first_name, last_name
	FROM indicadores.usuario usr
	JOIN indicadores.empresa empresa ON empresa.id = usr.empresa_id
	JOIN indicadores.nivel nivel ON nivel.id = empresa.nivel_id
	WHERE usr.id = $1
	`
	var user model.User
	rows, err := ur.connection.Query(query, id)
	if err != nil {
		return user, err
	}
	for rows.Next() {
		empresa := model.Empresa{}
		nivel := model.Nivel{}
		
		if err = rows.Scan(
			&user.ID, 
			&user.Usuario, 
			&user.Email, 
			&user.ManagerID, 
			&empresa.ID, 
			&empresa.Nome, 
			&nivel.ID, 
			&nivel.Nome,
			&user.FirstName,
			&user.LastName,
			); err != nil {
			return user, err
		}
		empresa.NivelID = nivel
		user.EmpresaID = &empresa

	}
	return user, nil
}

func (ur *UserRepository) Update(user *model.User) (*model.User, error) { return user, nil }

func (ur *UserRepository) Insert(user *model.User) error {
	query := `
		INSERT INTO indicadores.usuario(usuario, senha, email, empresa_id, manager_id, first_name, last_name)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		`
	hash, err := helpers.HashPassword(user.Senha)
	if err != nil {
		return err
	}
	if user.FirstName == "" || user.LastName == "" {
		return errors.New("first_name and last_name are required")
	}
	args := []interface{}{user.Usuario, hash, user.Email, user.EmpresaID.ID, user.ManagerID, user.FirstName, user.LastName}
	_, err = ur.connection.Exec(query, args...)
	return err
}

func (ur *UserRepository) DeleteGroup(grupo_id, user_id int) error { 
	query := `DELETE FROM indicadores.grupamento_usuario WHERE usuario_id = $1 AND grupamento_id = $2`
	_, err := ur.connection.Exec(query, user_id, grupo_id)
	return err
}

func (ur *UserRepository) FindByUsuarioSenha(usuario, senha string) (model.User, error) {
	query := `
	SELECT id, usuario, senha
	FROM indicadores.usuario
	where usuario = $1
	`
	var user model.User
	rows, err := ur.connection.Query(query, usuario)
	if err != nil {
		return user, err
	}
	for rows.Next() {
		if err = rows.Scan( &user.ID, &user.Usuario, &user.Senha); err != nil {
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
		if strings.Contains(err.Error(), "unicidade") {
			return errors.New("grupamento já adicionado")
		}
		return err
	}
	tx.Commit()
	return nil
}
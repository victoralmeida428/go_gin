package repository

import (
	"abramed_go/cmd/api/helpers"
	"abramed_go/cmd/api/model"
	"database/sql"
	"errors"
	"strings"
)

type UserRepository struct {
	connection *sql.DB
}

func (ur *UserRepository) FindAll() ([]model.User, error) {
	query := `
	SELECT id, usuario, senha, email, gerente_id
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
	SELECT id, usuario, senha, email, empresa_id
	FROM indicadores.usuario WHERE empresa_id = $1
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
	SELECT usr.id, usuario, email, empresa_id, empresa_id, empresa.nome, nivel.id, nivel.nome, usr.nome
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
			&user.Name,
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
		INSERT INTO indicadores.usuario(usuario, senha, email, empresa_id, gerente_id, nome)
		VALUES ($1, $2, $3, $4, $5, $6)
		`
	hash, err := helpers.HashPassword(user.Senha)
	if err != nil {
		return err
	}
	if user.Name == ""  {
		return errors.New("name is required")
	}
	args := []interface{}{user.Usuario, hash, user.Email, user.EmpresaID.ID, user.ManagerID, user.Name}
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
		if err = rows.Scan(&user.ID, &user.Usuario, &user.Senha); err != nil {
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

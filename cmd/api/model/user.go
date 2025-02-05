package model

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Empresa struct {
	ID      int    `json:"id"`
	Nome    string `json:"name"`
	NivelID Nivel  `json:"level"`
}

type Nivel struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

type User struct {
	ID          int          `json:"id"`
	Usuario     string       `json:"user,omitempty" binding:"required"`
	Name        string       `json:"name,omitempty" binding:"required"`
	Senha       string       `json:"password,omitempty" binding:"required"`
	Email       *string      `json:"email,omitempty"`
	EmpresaID   *Empresa     `json:"enterprise,omitempty"`
	ManagerID   *int         `json:"manager_id,omitempty"`
	Grupamentos []Grupamento `json:"groups,omitempty"`
}

var (
	JWT_SECRET_TOKEN = "JWT_SECRET_KEY"
)

func (user *User) GerarToken() (string, error) {
	secret := os.Getenv(JWT_SECRET_TOKEN)
	var (
		email string
	)
	if user.Email != nil {
		email = *user.Email
	}
	claims := jwt.MapClaims{
		"id":      user.ID,
		"usuario": user.Usuario,
		"email":   email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

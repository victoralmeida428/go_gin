package model

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type User struct {
	ID        int     `json:"id"`
	Usuario   string  `json:"usuario" binding:"required"`
	Senha     string  `json:"senha" binding:"required"`
	Email     *string `json:"email"`
	Empresa   *string `json:"empresa"`
	ManagerID *int    `json:"manager_id"`
}

var (
	JWT_SECRET_TOKEN = "JWT_SECRET_TOKEN"
)

func (user *User) GerarToken() (string, error) {
	secret := os.Getenv(JWT_SECRET_TOKEN)
	var email, empresa string
	var managerID int
	if user.Email != nil {
		email = *user.Email
	}
	if user.Empresa != nil {
		empresa = *user.Empresa
	}
	if user.ManagerID != nil {
		managerID = *user.ManagerID
	}
	claims := jwt.MapClaims{
		"id":         user.ID,
		"usuario":    user.Usuario,
		"email":      email,
		"empresa":    empresa,
		"manager_id": managerID,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

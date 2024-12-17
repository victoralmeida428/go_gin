package helpers

import (
	"abramed_go/model"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"os"
	"strings"
)

func GenerateSecretToken() string {
	// 32 bytes = 256 bits (é comum usar 256 bits para secret keys)
	const tokenSize = 32
	tokenBytes := make([]byte, tokenSize)

	// Preenche o slice com bytes aleatórios
	if _, err := rand.Read(tokenBytes); err != nil {
		log.Fatalf("Erro ao gerar o token: %v", err)
	}

	// Codifica os bytes em base64 para uma representação em string
	return base64.URLEncoding.EncodeToString(tokenBytes)
}

func VerifyToken(rawToken string) (*model.User, error) {
	secret := os.Getenv("JWT_SECRET_KEY")
	var user model.User
	token, err := jwt.Parse(RemoveBearerPrefix(rawToken), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid token")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return &user, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return &user, errors.New("Invalid token")
	}

	id, _ := claims["id"].(float64)

	// Converte "manager_id" para *int
	var managerID *int
	if val, ok := claims["manager_id"].(float64); ok {
		intVal := int(val)
		managerID = &intVal
	}

	// Converte "empresa" e "email" para *string
	var empresa *string
	if val, ok := claims["empresa"].(string); ok {
		empresa = &val
	}

	var email *string
	if val, ok := claims["email"].(string); ok {
		email = &val
	}

	return &model.User{
		ID:        int(id),
		Usuario:   claims["usuario"].(string),
		Empresa:   empresa,
		Email:     email,
		ManagerID: managerID,
	}, nil

}

func RemoveBearerPrefix(token string) string {
	if strings.HasPrefix(token, "Bearer ") {
		token = strings.TrimPrefix(token, "Bearer ")
	}
	return token
}

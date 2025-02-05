package helpers

import (
	"abramed_go/cmd/api/model"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"log"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
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
			return nil, errors.New("invalid token")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return &user, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return &user, errors.New("invalid token")
	}

	id, _ := claims["id"].(float64)

	var email *string
	if val, ok := claims["email"].(string); ok {
		email = &val
	}

	return &model.User{
		ID:      int(id),
		Usuario: claims["usuario"].(string),
		Email:   email,
	}, nil

}

func RemoveBearerPrefix(token string) string {
	return strings.TrimPrefix(token, "Bearer ")
}

package helpers

import (
	"abramed_go/model"
	"testing"
)

func TestVerify(t *testing.T) {
	var (
		empresa = "empresa"
		email   = "email@email.com"
	)
	user := model.User{ID: 1, Usuario: "victor", Senha: "teste", Empresa: &empresa, Email: &email}
	token, err := user.GerarToken()
	if err != nil {
		t.Error(err)
	}
	u, err := VerifyToken(token)
	if err != nil {
		t.Error(err)
	}
	if u.ID != user.ID {
		t.Error("não retornou o usuário correto")
	}
}

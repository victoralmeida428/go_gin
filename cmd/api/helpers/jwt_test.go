package helpers

import (
	"abramed_go/cmd/api/model"
	"testing"
)

func TestVerify(t *testing.T) {
	var (
		empresa = model.Empresa{ID: 1, Nome: "empresa"}
		email   = "email@email.com"
	)
	user := model.User{ID: 1, Usuario: "victor", Senha: "teste", EmpresaID: &empresa, Email: &email}
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

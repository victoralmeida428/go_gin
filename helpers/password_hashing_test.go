package helpers

import (
	"strings"
	"testing"
)

func FuzzHashing(f *testing.F) {
	f.Add("testandosenha")
	f.Fuzz(func(t *testing.T, senha string) {
		if len([]byte(senha)) > 72 {
			return
		}
		hash, err := HashPassword(senha)
		if err != nil {
			t.Error(err)
		}
		if !strings.HasPrefix(hash, "$2a$10$") {
			t.Error("Hashes don't match with the bcrypt with 10 cost")
		}
		if !CheckPasswordHash(senha, hash) {
			t.Error("Hashes don't match with the password")
		}
	})
}

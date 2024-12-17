package model

type Formulario struct {
	ID        int    `json:"id"`
	Nome      string `json:"nome" binding:"required"`
	Descricao string `json:"descricao"`
	Ativo     bool   `json:"ativo" binding:"required"`
}

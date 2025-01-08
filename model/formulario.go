package model

import "time"

type Formulario struct {
	ID        int    `json:"id"`
	Nome      string `json:"name" binding:"required"`
	Descricao string `json:"description"`
	Ativo     bool   `json:"active" binding:"required"`
}

type VersaoFormulario struct {
	ID         int        `json:"id"`
	Formulario Formulario `json:"form"`
	Versao     int        `json:"version" binding:"required"`
	CriadoEm   time.Time  `json:"created_at"`
}

type Metodo struct {
	ID   int    `json:"id"`
	Nome string `json:"name"`
}
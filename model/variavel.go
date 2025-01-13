package model

import "time"

type Variavel struct {
	ID           int          `json:"id,omitempty"`
	TipoVariavel *TipoVariavel `json:"data_type,omitempty" binding:"required"`
	PerguntaId   *int         `json:"question_id,omitempty"`
	PossuiItem   bool         `json:"has_item,omitempty"`
	Obrigatorio  bool         `json:"required,omitempty"`
	Texto        string       `json:"text,omitempty" binding:"required"`
}

type ViewVariavel struct {
	ID           int          `json:"id,omitempty"`
	TipoVariavel TipoVariavel `json:"data_type,omitempty" binding:"required"`
	Pergunta     string      `json:"question,omitempty"`
	Item         *string      `json:"item,omitempty"`
	PossuiItem   bool         `json:"has_item,omitempty"`
	Obrigatorio  bool         `json:"required,omitempty"`
}

type TipoVariavel struct {
	ID   int    `json:"id"`
	Nome string `json:"nome" binding:"required"`
}

type VersaoVariavel struct {
	ID           int        `json:"id"`
	FormularioId Formulario `json:"form_id,omitempty" binding:"required"`
	Versao       int        `json:"version,omitempty" binding:"required"`
	CriadoEm     time.Time  `json:"created_at,omitempty" binding:"required"`
}

type Grupamento struct {
	ID   int    `json:"id"`
	Nome string `json:"name" binding:"required"`
}

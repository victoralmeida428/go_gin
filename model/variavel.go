package model

import "time"

type Variavel struct {
	ID           int          `json:"id"`
	TipoVariavel TipoVariavel `json:"data_type_id" binding:"required"`
	PerguntaId   *int         `json:"question_id,omitempty"`
	PossuiItem   bool         `json:"has_item"`
	Obrigatorio  bool         `json:"required"`
	Texto        string       `json:"text" binding:"required"`
}

type ViewVariavel struct {
	ID           int          `json:"id"`
	TipoVariavel TipoVariavel `json:"data_type_id" binding:"required"`
	Pergunta     string      `json:"question,omitempty"`
	Item         *string      `json:"item,omitempty"`
	PossuiItem   bool         `json:"has_item"`
	Obrigatorio  bool         `json:"required"`
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

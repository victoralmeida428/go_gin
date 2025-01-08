package model

import "time"

type Variavel struct {
	ID           int          `json:"id"`
	TipoVariavel TipoVariavel `json:"tipo_variavel_id" binding:"required"`
	PerguntaId   *int         `json:"pergunta_id,omitempty"`
	PossuiItem   bool         `json:"possui_item"`
	Obrigatorio  bool         `json:"obrigatorio"`
	Texto        string       `json:"texto" binding:"required"`
}

type ViewVariavel struct {
	ID           int          `json:"id"`
	TipoVariavel TipoVariavel `json:"tipo_variavel_id" binding:"required"`
	Pergunta     string      `json:"pergunta,omitempty"`
	Item         *string      `json:"item,omitempty"`
	PossuiItem   bool         `json:"possui_item"`
	Obrigatorio  bool         `json:"obrigatorio"`
}

type TipoVariavel struct {
	ID   int    `json:"id"`
	Nome string `json:"nome" binding:"required"`
}

type VersaoVariavel struct {
	ID           int        `json:"id"`
	FormularioId Formulario `json:"formulario_id,omitempty" binding:"required"`
	Versao       int        `json:"versao,omitempty" binding:"required"`
	CriadoEm     time.Time  `json:"criado_em,omitempty" binding:"required"`
}

type Grupamento struct {
	ID   int    `json:"id"`
	Nome string `json:"nome" binding:"required"`
}

package model

import "time"

type Variavel struct {
	ID           int    `json:"id"`
	GrupamentoId *int   `json:"grupamento_id,omitempty"`
	TipoVariavel int    `json:"tipo_variavel_id" binding:"required"`
	PerguntaId   *int   `json:"pergunta_id,omitempty"`
	PossuiItem   bool   `json:"possui_item"`
	Obrigatorio  bool   `json:"obrigatorio"`
	Texto        string `json:"texto" binding:"required"`
}

type TipoVariavel struct {
	ID   int    `json:"id"`
	Nome string `json:"nome" binding:"required"`
}

type VersaoVariavel struct {
	ID           int       `json:"id"`
	FormularioId int       `json:"formulario_id" binding:"required"`
	Versao       int       `json:"versao" binding:"required"`
	CriadoEm     time.Time `json:"criado_em" binding:"required"`
}

type Grupamento struct {
	ID   int    `json:"id"`
	Nome string `json:"nome" binding:"required"`
}

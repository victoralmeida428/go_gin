package dto_agendamento

import "time"

type CreateInput struct {
	Year int `json:"year" binding:"required"`
	Versao int `json:"version_form_id" binding:"required"`
	Periodicidade  int `json:"periodicity_id" binding:"required"`
	Metodo int `json:"method_id" binding:"required"`
	Inicio time.Time
	Ativo bool
}


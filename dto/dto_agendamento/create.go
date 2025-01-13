package dto_agendamento

type CreateInput struct {
	Versao int `json:"version_form_id" binding:"required"`
	Periodicidade int `json:"periodicity_id" binding:"required"`
	Inicio string `json:"start" binding:"required"`
	Ativo bool `json:"active" binding:"required"`
	Metodo int `json:"method_id" binding:"required"`

}
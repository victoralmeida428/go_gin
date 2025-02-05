package dto_respostas

type ResponderInput struct {
	Variavel    int     `json:"variable_id" binding:"required"`
	Agendamento int     `json:"schedule_id" binding:"required"`
	Valor       float64 `json:"value"`
}

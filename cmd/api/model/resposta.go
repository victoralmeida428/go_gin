package model

import "time"

type RespostaNum struct {
	ID       int       `json:"id,omitempty"`
	Valor    float64   `json:"value"`
	Versao   int16     `json:"version"`
	CriadoEm time.Time `json:"created_at"`
}

type Resposta struct {
	ID          int           `json:"id,omitempty"`
	Variavel    Variavel      `json:"variable,omitempty"`
	Agendamento *Agendamento   `json:"schedule,omitempty"`
	Usuario     User          `json:"user,omitempty"`
	Respostas   []RespostaNum `json:"answer,omitempty"`
}

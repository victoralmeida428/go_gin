package model

import "time"

type Periodicidade struct {
	ID        int    `json:"id"`
	Nome      string `json:"name"`
	Intervalo string `json:"interval,omitempty"`
}

type Agendamento struct {
	ID               int              `json:"id"`
	VersaoFormulario VersaoFormulario `json:"version_form,omitempty"`
	Periodicidade    Periodicidade    `json:"periodicity"`
	Inicio           time.Time        `json:"start"`
	Proximo          time.Time        `json:"next"`
	Ativo            bool             `json:"active"`
	Metodo           Metodo           `json:"method"`
	User             User             `json:"user"`
}

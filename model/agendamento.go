package model

import "time"

type Periodicidade struct {
	ID        int    `json:"id,omitempty"`
	Nome      string `json:"name,omitempty"`
	Intervalo string `json:"interval,omitempty"`
}

type Agendamento struct {
	ID               int              `json:"id,omitempty"`
	VersaoFormulario VersaoFormulario `json:"version_form,omitempty"`
	Periodicidade    Periodicidade    `json:"periodicity,omitempty"`
	Inicio           time.Time        `json:"start,omitempty"`
	Proximo          time.Time        `json:"next,omitempty"`
	Ativo            bool             `json:"active,omitempty"`
	Metodo           Metodo           `json:"method,omitempty"`
	User             User             `json:"user,omitempty"`
}

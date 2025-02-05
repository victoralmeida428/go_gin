package model

import (
	"abramed_go/cmd/api/helpers/status"
	"time"
)

type Periodicidade struct {
	ID        int    `json:"id,omitempty"`
	Nome      string `json:"name,omitempty"`
	Intervalo string `json:"interval,omitempty"`
}

type Agendamento struct {
	ID               int               `json:"id,omitempty"`
	VersaoFormulario VersaoFormulario  `json:"version_form,omitempty"`
	Periodicidade    Periodicidade     `json:"periodicity,omitempty"`
	Inicio           time.Time         `json:"start,omitempty"`
	End              time.Time         `json:"end,omitempty"`
	Ativo            bool              `json:"active"`
	Metodo           Metodo            `json:"method,omitempty"`
	User             User              `json:"user,omitempty"`
	Deadline         time.Time         `json:"deadline"`
	Week             *int              `json:"week"`
	Year             *int              `json:"year"`
	Status           status.StatusForm `json:"status_id"`
}

// GetWeek determina o número da semana do ano e o ano correspondente,
// seguindo o calendário epidemiológico do Ministério da Saúde do Brasil.
//
// O calendário epidemiológico define a primeira semana epidemiológica do ano
// como aquela que contém pelo menos quatro dias do ano novo, e cada semana
// subsequente possui exatamente sete dias. Este calendário é frequentemente
// utilizado para a vigilância e análise epidemiológica.
//
// Parâmetros:
//
//	data (time.Time): A data para a qual o número da semana epidemiológica será calculado.
//
// Retorno:
//
//	year (int): O ano correspondente à semana epidemiológica. Pode diferir do ano
//	            do parâmetro 'data' se a data estiver em uma semana que pertence ao
//	            ano epidemiológico anterior ou seguinte.
//	week (int): O número da semana epidemiológica no ano correspondente.
//
// Exemplo de uso:
//
//	data := time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC)
//	year, week := getWeek(data)
//	fmt.Printf("Ano: %d, Semana: %d\n", year, week)
//
// Este cálculo é fundamental para análises de dados epidemiológicos, pois permite
// que informações sejam organizadas de forma consistente ao longo das semanas epidemiológicas.
func (agenda *Agendamento) GetWeek() (year int, week int) {
	year, week = agenda.Inicio.ISOWeek()
	if week == 52 {
		if 31-agenda.Inicio.Day() >= 3 && agenda.Inicio.Day() > 20 {
			week = 53
			return
		} else {
			week = 1
			year++
		}
	} else {
		year, week = agenda.Inicio.Add(24 * 2 * time.Hour).ISOWeek()
	}
	return

}

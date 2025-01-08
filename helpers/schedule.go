package helpers

import (
	"errors"
	"fmt"
	"time"
)

type Schedule uint8
const (
	Diario Schedule = iota
	Semanal
	Mensal
	Trimestral
	Semestral
	Anual
)

func ParseCron(agenda Schedule, date time.Time)(string, error) {
	day := date.Day()
	week := int(date.Weekday())

	switch agenda {
	case Diario:

		return "0 0 * * *", nil
	case Semanal:
		return fmt.Sprintf("0 0 * * %d", week), nil
	case Mensal:
		return fmt.Sprintf("0 0 %d * *", day), nil
	case Trimestral:
		return fmt.Sprintf("0 0 %d 1-12/3 *", day), nil
	case Semestral:
		return fmt.Sprintf("0 0 %d 1-12/6 *", day), nil
	case Anual:
		return fmt.Sprintf("0 0 %d 1-12/12 *", day), nil
	}
	return "", errors.New("invalid schedule")
}
package status

import (
	"time"
)

type StatusForm int

func (status StatusForm) ToString() string {
	switch status {
	case WAITING:
		return "REPORTAR"
	case PENDING:
		return "PENDENTE"
	case COMPLETE:
		return "COMPLETO"
	case PARTIAL:
		return "PARCIAL"
	case PARTIALBACKDATE:
		return "PARCIAL R."
	case COMPLETEBACKDATE:
		return "COMPLETO R."
	default:
		return "DESCONHECIDO"
	}
}

const (
	WAITING StatusForm = iota
	PENDING
	PARTIAL
	PARTIALBACKDATE
	COMPLETE
	COMPLETEBACKDATE
)

func GetStatusForm(totalRequired, totalOptional, required, optional int, deadline, report_opt, report_required time.Time) StatusForm {
	answerBackdate := report_opt.After(deadline) && report_required.After(deadline)
	requiredBackdate := report_required.After(deadline)

	if required == totalRequired && optional == totalOptional && answerBackdate {
		return COMPLETEBACKDATE

	} else if required == totalRequired && optional == totalOptional {
		return COMPLETE

	} else if required == totalRequired && optional < totalOptional && requiredBackdate {
		return PARTIALBACKDATE

	} else if required == totalRequired && optional < totalOptional {
		return PARTIAL

	} else if !answerBackdate && !time.Now().After(deadline) {
		return WAITING
	}
	return PENDING
}

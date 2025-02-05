package controller

import (
	"abramed_go/cmd/api/dto/dto_agendamento"
	"abramed_go/cmd/api/helpers"
	"abramed_go/cmd/api/helpers/status"
	"abramed_go/cmd/api/model"
	"abramed_go/cmd/api/repository"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type agendamentoController struct {
	repo *repository.Repository
}

// @Summary List Schedules
// @Description List all schedules
// @Tags schedules
// @Accept json
// @Security BearerAuth
// @Produce json
// @Success 200 {object} []model.Agendamento
// @Router /schedule [get]
func (c *agendamentoController) ListAll(ctx *gin.Context) {
	userHeader, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}
	user, err := c.repo.User.FindById(userHeader.(*model.User).ID)
	user.Senha = ""
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	stat, err := strconv.Atoi(ctx.Query("status"))
	var statusForm *status.StatusForm
	if err == nil {
		statForm := status.StatusForm(stat)
		statusForm = &statForm
	}

	agendamentos, err := c.repo.Agendamento.FindAll(&user, statusForm)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	type outputType struct {
		model.Agendamento
		Week   int    `json:"week"`
		Year   int    `json:"year"`
		Status string `json:"status"`
	}

	output := make([]outputType, len(agendamentos))
	for i := range agendamentos {
		var row outputType
		row.Agendamento = agendamentos[i]
		year, week := agendamentos[i].GetWeek()
		row.Week = week
		row.Year = year
		row.Status = agendamentos[i].Status.ToString()

		output[i] = row

	}
	ctx.JSON(200, output)
}

func (c *agendamentoController) createSchedule(userId int, input *dto_agendamento.CreateInput, start time.Time) error {
	input.Inicio = start
	input.Ativo = true

	return c.repo.Agendamento.Insert(userId, input)

}

func (c *agendamentoController) Create(ctx *gin.Context) {
	userHeader, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}
	var input dto_agendamento.CreateInput
	if err := ctx.ShouldBindBodyWithJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var err error
	start := helpers.GetFirstSunday(input.Year)
	newYear, _ := helpers.GetWeek(start)
	for newYear == input.Year {
		//cadastrar
		fmt.Println(newYear)
		err = c.createSchedule(userHeader.(*model.User).ID, &input, start)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		start = start.Add(7 * helpers.Day)
		newYear, _ = helpers.GetWeek(start)
		if newYear > input.Year {
			break
		}
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Scheduling created with sucess"})

}

func (c *agendamentoController) ListPeriodicity(ctx *gin.Context) {
	agendamentos, err := c.repo.Agendamento.ListPeriodicity()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, agendamentos)

}

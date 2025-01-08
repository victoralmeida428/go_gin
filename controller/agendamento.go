package controller

import (
	"abramed_go/model"
	"abramed_go/repository"
	"net/http"

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

	agendamentos, err := c.repo.Agendamento.FindAll(&user)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, agendamentos)
}
package controller

import (
	"abramed_go/cmd/api/dto/dto_respostas"
	"abramed_go/cmd/api/helpers"
	"abramed_go/cmd/api/model"
	"abramed_go/cmd/api/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type respostaController struct {
	repo *repository.Repository
}

func (c *respostaController) Responder(ctx *gin.Context) {
	var inputs []dto_respostas.ResponderInput
	if err := ctx.ShouldBindBodyWithJSON(&inputs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userHeader, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}
	for _, input := range inputs {
		if err := c.repo.Resposta.Responder(input, userHeader.(*model.User).ID); err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "respondido com sucesso"})

}

func (c *respostaController) ListRepostas(ctx *gin.Context) {
	idSchedule, err := helpers.QueryInt(ctx, "schedule_id")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	output, err := c.repo.Resposta.ListAll(idSchedule)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, output)
}

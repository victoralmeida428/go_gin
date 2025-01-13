package controller

import (
	"abramed_go/dto/dto_respostas"
	"abramed_go/helpers"
	"abramed_go/model"
	"abramed_go/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)


type respostaController struct {
	repo *repository.Repository
}

func (c *respostaController) Responder(ctx *gin.Context){
    var input []dto_respostas.ResponderInput
    if err := ctx.ShouldBindBodyWithJSON(&input); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userHeader, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

    msg := make(map[int]string)
    for i := range input {
        if err := c.repo.Resposta.Responder(input[i], userHeader.(*model.User).ID); err != nil {
            ctx.JSON(http.StatusInternalServerError, err.Error())
            return
        }
        msg[input[i].Variavel] = "Pergunta respondida com sucesso"
    }
    ctx.JSON(http.StatusCreated, msg)

}

func (c *respostaController) ListRepostas(ctx *gin.Context) {
    idSchedule, err :=  helpers.QueryInt(ctx, "schedule_id")
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
        return
    }
    output, err := c.repo.Resposta.ListAll(idSchedule)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    }
    ctx.JSON(http.StatusOK, output)
}


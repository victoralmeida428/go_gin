package controller

import (
	"abramed_go/model"
	"abramed_go/repository"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type variavelController struct {
	repo *repository.Repository
}

// @Summary Listar tipos
// @Description Listar todos tipos de variáveis
// @Tags variaveis
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} []model.TipoVariavel
// @Router /api/variaveis/tipo [get]
func (c *variavelController) ListTipos(ctx *gin.Context){
	tipos, err := c.repo.Variavel.ListTipos()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tipos)

}


// @Summary Criar grupamento
// @Description Criar um grupamento
// @Tags variaveis
// @Accept json
// @Produce json
// @Param user body model.Grupamento true  "Grupamento"
// @Security BearerAuth
// @Success 200 {object} createResponse
// @Router /api/variaveis/grupamento [put]
func (c *variavelController) CreateGrupo(ctx *gin.Context){
	
	var input model.Grupamento
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := c.repo.Variavel.CreateGrupamento(&input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message":fmt.Sprintf("Grupamento %s criado com sucesso", input.Nome)})
}


// @Summary Listar Grupamentos
// @Description Listar todos os Grupamentos
// @Tags variaveis
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} []model.Grupamento
// @Router /api/variaveis/grupamento [get]
func (c *variavelController) ListGrupo(ctx *gin.Context){
	
	grupos, err := c.repo.Variavel.ListGrupamentos()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	ctx.JSON(http.StatusOK, grupos)
}

// @Summary Criar Variáveis
// @Description Criar somente a variável
// @Tags variaveis
// @Accept json
// @Produce json
// @Param user body model.Variavel true  "Variavel"
// @Security BearerAuth
// @Success 200 {object} createResponse
// @Router /api/variaveis [put]
func (c *variavelController) CreateVariavel(ctx *gin.Context){
	var input model.Variavel
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.repo.Variavel.Insert(&input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message":"Variavel criado com sucesso"})

}

// @Summary Listar Variaveis
// @Description Listar todos as Variaveis
// @Tags variaveis
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} []model.Variavel
// @Router /api/variaveis [get]
func (c *variavelController) ListVariavel(ctx *gin.Context) {
	variaveis, err := c.repo.Variavel.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, variaveis)
}


// @Summary Listar Grupamentos
// @Description Listar todos os grupamentos existentes
// @Tags variaveis
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} []model.Grupamento
// @Router /api/variaveis/grupamento [get]
func (c *variavelController) ListGrupamentos(ctx *gin.Context) {
	userHeader, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	user, err := c.repo.User.FindById(userHeader.(*model.User).ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	variaveis, err := c.repo.Variavel.ListGrupamentosByUser(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, variaveis)
}


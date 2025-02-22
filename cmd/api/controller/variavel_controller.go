package controller

import (
	"abramed_go/cmd/api/helpers"
	"abramed_go/cmd/api/model"
	"abramed_go/cmd/api/repository"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type variavelController struct {
	repo *repository.Repository
}

// @Summary List types
// @Description List all variables types
// @Tags variables
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} []model.TipoVariavel
// @Router /variable/tipo [get]
func (c *variavelController) ListTipos(ctx *gin.Context) {
	tipos, err := c.repo.Variavel.ListTipos()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tipos)

}

// @Summary Create groups
// @Tags variables
// @Accept json
// @Produce json
// @Param user body model.Grupamento true  "Groups"
// @Security BearerAuth
// @Success 200 {object} createResponse
// @Router /variable/groups [put]
func (c *variavelController) CreateGrupo(ctx *gin.Context) {

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
	ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Grupamento %s criado com sucesso", input.Nome)})
}

// @Summary List Groups
// @Description List all groups
// @Tags variaveis
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} []model.Grupamento
// @Router /variable/groups [get]
func (c *variavelController) ListGrupo(ctx *gin.Context) {

	grupos, err := c.repo.Variavel.ListGrupamentos()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, grupos)
}

// @Summary Create variable
// @Description Create only one variable
// @Tags variables
// @Accept json
// @Produce json
// @Param user body model.Variavel true  "Variable"
// @Security BearerAuth
// @Success 200 {object} createResponse
// @Router /variable [put]
func (c *variavelController) CreateVariavel(ctx *gin.Context) {
	var input model.Variavel
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.repo.Variavel.Insert(&input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Variavel criado com sucesso"})

}

// @Summary List variables
// @Description List all variables
// @Tags variables
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} []model.Variavel
// @Router /variable [get]
func (c *variavelController) ListVariavel(ctx *gin.Context) {
	formId, err := helpers.QueryInt(ctx, "id")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	variaveis, err := c.repo.Variavel.FindAll(formId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, variaveis)
}

// @Summary List Groups by user
// @Description List all groups by user
// @Tags variables
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} []model.Grupamento
// @Router /variable/groups [get]
// func (c *variavelController) ListGrupamentos(ctx *gin.Context) {
// 	userHeader, exists := ctx.Get("user")
// 	if !exists {
// 		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
// 		return
// 	}

// 	user, err := c.repo.User.FindById(userHeader.(*model.User).ID)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	variaveis, err := c.repo.Variavel.ListGrupamentosByUser(&user)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, err)
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, variaveis)
// }

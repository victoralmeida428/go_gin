package controller

import (
	"abramed_go/model"
	"abramed_go/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

type formController struct {
	formRepo repository.IRepository[model.Formulario]
}

// Lista todos os formulários
// @Summary Formulário
// @Description Retornar todos os formulários ativos
// @Tags formulario
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} []model.Formulario
// @Router /api/formulario [get]
func (fc *formController) List(ctx *gin.Context) {
	forms, err := fc.formRepo.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, forms)
}

type createResponse struct {
	Message string `form:"message" json:"message" binding:"required"`
}

// Cria um formulario
// @Summary Criar Formulário
// @Description Retornar todos os formulários ativos
// @Tags formulario
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param form body model.Formulario true "Formulário"
// @Success 200 {object} createResponse
// @Router /api/formulario [put]
func (fc *formController) Create(ctx *gin.Context) {
	var form model.Formulario
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := fc.formRepo.Insert(&form)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, createResponse{Message: "Formulario criado com sucesso"})
}

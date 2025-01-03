package controller

import (
	"abramed_go/model"
	"abramed_go/repository"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type formController struct {
	repo *repository.Repository
}

// Lista todos os formulários
// @Summary Listar Formulário
// @Description Retornar todos os formulários ativos
// @Tags formulario
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} []model.Formulario
// @Router /api/formulario [get]
func (fc *formController) List(ctx *gin.Context) {
	forms, err := fc.repo.Formulario.FindAll()
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
	err := fc.repo.Formulario.Insert(&form)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, createResponse{Message: "Formulario criado com sucesso"})
}

// Deleta um formulario
// @Summary Deletar Formulário
// @Description Deleta o usuário
// @Tags formulario
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param form body model.Formulario true "Formulário"
// @Success 200 {object} createResponse
// @Router /api/formulario/{id} [delete]
func (fc *formController) Delete(ctx *gin.Context) {
	id, exists := ctx.Params.Get("id")
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id não informado"})
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	form, _ := fc.repo.Formulario.FindById(idInt)

	
	err = fc.repo.Formulario.Delete(idInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, createResponse{Message: fmt.Sprintf("Formulario: %s Deletado com sucesso", form.Nome)})
}


// Atualiza um formulario
// @Summary Atualizar Formulário
// @Description Atualiza um usuário
// @Tags formulario
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param form body model.Formulario true "Formulário"
// @Success 200 {object} createResponse
// @Router /api/formulario [patch]
func (fc *formController) Update(ctx *gin.Context) {
	type input struct {
		ID int `json:"id" binding:"required"`
		Ativo *bool `json:"ativo,omitempty"`
		Descricao *string `json:"descricao,omitempty"`
		Nome *string `json:"nome,omitempty"`
	}
	var formInput input
	if err := ctx.ShouldBindJSON(&formInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	form, err := fc.repo.Formulario.FindById(formInput.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if formInput.Ativo != nil {
		form.Ativo = *formInput.Ativo
	}
	if formInput.Nome != nil && *formInput.Nome != "" {
		form.Nome = *formInput.Nome
	}
	if formInput.Descricao != nil && *formInput.Descricao != ""{
		form.Descricao = *formInput.Descricao
	}

	_, err = fc.repo.Formulario.Update(&form)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, createResponse{Message: "Formulario Atualizado com sucesso"})
}
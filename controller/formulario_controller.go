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

// @Summary List Forms
// @Description List all forms
// @Tags form
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} []model.Formulario
// @Router /form [get]
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


// @Summary Create Form
// @Description Create a form
// @Tags form
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param form body model.Formulario true "Input"
// @Success 200 {object} createResponse
// @Router /form [put]
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


// @Summary Delete Form
// @Description Delete Form
// @Tags form
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param form body model.Formulario true "form"
// @Success 200 {object} createResponse
// @Router /form/{id} [delete]
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


// @Summary Update Form
// @Tags form
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param form body model.Formulario true "Input"
// @Success 200 {object} createResponse
// @Router /form [patch]
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
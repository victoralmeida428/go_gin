package controller

import (
	"abramed_go/cmd/api/model"
	"abramed_go/cmd/api/repository"
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
func (c *formController) List(ctx *gin.Context) {
	forms, err := c.repo.Formulario.Query(0)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, forms)
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
func (c *formController) Create(ctx *gin.Context) {
	var form model.Formulario
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := c.repo.Formulario.Insert(&form)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Formulario criado com sucesso"})
}

func (c *formController) FindById(ctx *gin.Context) {
	id, exists := ctx.Params.Get("id")
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Id required"})
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	form, err := c.repo.Formulario.Query(idInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, form)

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
func (c *formController) Delete(ctx *gin.Context) {
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

	form, _ := c.repo.Formulario.FindById(idInt)

	err = c.repo.Formulario.Delete(idInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Formulario: %s Deletado com sucesso", form.Formulario.Nome)})
}

// @Summary Update Form
// @Tags form
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param form body model.Formulario true "Input"
// @Success 200 {object} createResponse
// @Router /form [patch]
func (c *formController) Update(ctx *gin.Context) {
	type input struct {
		ID        int     `json:"id" binding:"required"`
		Ativo     *bool   `json:"ativo,omitempty"`
		Descricao *string `json:"descricao,omitempty"`
		Nome      *string `json:"nome,omitempty"`
	}
	var formInput input
	if err := ctx.ShouldBindJSON(&formInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	form, err := c.repo.Formulario.FindById(formInput.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if formInput.Ativo != nil {
		form.Formulario.Ativo = *formInput.Ativo
	}
	if formInput.Nome != nil && *formInput.Nome != "" {
		form.Formulario.Nome = *formInput.Nome
	}
	if formInput.Descricao != nil && *formInput.Descricao != "" {
		form.Formulario.Descricao = *formInput.Descricao
	}

	_, err = c.repo.Formulario.Update(&form.Formulario)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Formulario Atualizado com sucesso"})
}

func (c *formController) ListMethods(ctx *gin.Context) {
	metodos, err := c.repo.Formulario.ListMethods()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, metodos)
}

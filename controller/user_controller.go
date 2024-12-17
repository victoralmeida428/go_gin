package controller

import (
	"abramed_go/model"
	"abramed_go/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userController struct {
	userUsecase repository.IUserRepository
}

// GetUser
// @Summary Criar Usuário
// @Description Criar um novo usuário
// @Tags user
// @Accept json
// @Produce json
// @Param user body model.User true  "Usuário"
// @Success 200 {object} model.User
// @Router /api/user/create [put]
func (uc *userController) CreateUser(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := uc.userUsecase.Insert(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, user)

}

// GetUser
// @Summary Usuário
// @Description Retornar os dados do usuário logado
// @Tags user
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} model.User
// @Router /api/user [get]
func (uc *userController) GetUser(ctx *gin.Context) {
	userHeader, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	user, err := uc.userUsecase.FindById(userHeader.(*model.User).ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

type loginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type loginResponse struct {
	Token string `json:"token"`
}

// Fazer login
// @Summary Login
// @Description Pegar token JWT
// @Tags user
// @Accept json
// @Produce json
// @Param login body loginRequest true "Dados de Login"
// @Success 200 {object} loginResponse
// @Router /api/user/login [post]
func (uc *userController) Login(ctx *gin.Context) {
	var user loginRequest
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userAuth, err := uc.userUsecase.FindByUsuarioSenha(user.Username, user.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := userAuth.GerarToken()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, loginResponse{Token: token})
}

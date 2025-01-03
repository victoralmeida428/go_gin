package controller

import (
	"abramed_go/model"
	"abramed_go/repository"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userController struct {
	repo *repository.Repository
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
	if err := uc.repo.User.Insert(&user); err != nil {
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

	user, err := uc.repo.User.FindById(userHeader.(*model.User).ID)
	user.Senha=""

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	members, err := uc.repo.User.FindAllMembers(user.ID)
	for i := range members {
		members[i].Senha = ""
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	groups, err := uc.repo.Variavel.ListGrupamentosByUser(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	type membersS struct {
		model.User
		Grupos []model.Grupamento `json:"grupos"`
	} 
	membersOut := make([]membersS, 0)

	for i := range members {
		groups, err := uc.repo.Variavel.ListGrupamentosByUser(&members[i])
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		membersOut = append(membersOut, membersS{members[i], groups})
	}

	if user.ManagerID == nil {
		ctx.JSON(http.StatusOK, gin.H{"user": user, "members":membersOut})
		return
	} 

	ctx.JSON(http.StatusOK, gin.H{"user": membersS{user, groups}, })
	
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

	userAuth, err := uc.repo.User.FindByUsuarioSenha(user.Username, user.Password)
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

// @Summary Adicionar Grupamento
// @Description Fazer ligação do usuário com o grupamento
// @Tags user
// @Accept json
// @Produce json
// @Param login body loginRequest true "Dados de Login"
// @Success 200 {object} loginResponse
// @Router /api/user/grupamento [post]
func (uc *userController) AddGrupamento(ctx *gin.Context) {
	var input struct {
		UserID       []int `json:"user_id" binding:"required"`
		GrupamentoID int `json:"grupamento_id" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	grupo, err := uc.repo.Variavel.FindGrupamentoById(input.GrupamentoID)
	fmt.Printf("%+v\n",grupo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	for _, userId := range input.UserID {
		user, err := uc.repo.User.FindById(userId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
	
	
		err = uc.repo.User.AddGrupamento(&user, grupo)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
	}
	
	ctx.JSON(http.StatusOK, gin.H{"message":fmt.Sprintf("Usuários adicionado ao grupo %s", grupo.Nome)})
}

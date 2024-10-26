package handler

import (
	"net/http"
	"server-pulsa-app/internal/entity/dto"
	"server-pulsa-app/internal/logger"
	"server-pulsa-app/internal/usecase"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authUsecase usecase.AuthUseCase
	rg          *gin.RouterGroup
	log         *logger.Logger
}

func (a *AuthController) loginHandler(ctx *gin.Context) {
	var payload dto.AuthRequestDto

	a.log.Info("Starting to login a user in the handler layer", nil)

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		a.log.Error("Invalid payload for login", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	a.log.Info("Starting login", nil)
	token, err := a.authUsecase.Login(payload)
	if err != nil {
		a.log.Error("Failed to authenticate user: ", err)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	a.log.Info("User has been authenticated successfully", nil)
	ctx.JSON(http.StatusOK, token)
}

func (a *AuthController) registerHandler(ctx *gin.Context) {
	var payload dto.AuthRequestDto

	a.log.Info("Starting to register a new user in the handler layer", nil)
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		a.log.Error("Invalid payload for register", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	a.log.Info("Starting to register new user", nil)
	user, err := a.authUsecase.Register(payload)
	if err != nil {
		a.log.Error("Failed to register user: ", err)
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	a.log.Info("User has been registered successfully", nil)
	ctx.JSON(http.StatusCreated, user)
}

func (a *AuthController) Route() {
	a.rg.POST("/auth/login", a.loginHandler)
	a.rg.POST("/auth/register", a.registerHandler)
}

func NewAuthController(authUc usecase.AuthUseCase, rg *gin.RouterGroup, log *logger.Logger) *AuthController {
	return &AuthController{authUsecase: authUc, rg: rg, log: log}
}

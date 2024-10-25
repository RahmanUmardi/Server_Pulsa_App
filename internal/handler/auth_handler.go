package handler

import (
	"net/http"
	"server-pulsa-app/internal/entity/dto"
	"server-pulsa-app/internal/usecase"

	"github.com/gin-gonic/gin"
)

// var logrus = logger.GetLogger()

type AuthController struct {
	authUsecase usecase.AuthUseCase
	rg          *gin.RouterGroup
}

func (a *AuthController) loginHandler(ctx *gin.Context) {
	var payload dto.AuthRequestDto

	// logrus.Info("Starting to login a user in the handler layer")

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		// logrus.Errorf("Invalid payload for login: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// logrus.Info("Starting login")
	token, err := a.authUsecase.Login(payload)
	if err != nil {
		// logrus.Error("Failed to authenticate user: ", err)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// logrus.Info("User has been authenticated successfully")
	ctx.JSON(http.StatusOK, token)
}

func (a *AuthController) registerHandler(ctx *gin.Context) {
	var payload dto.AuthRequestDto

	// logrus.Info("Starting to register a new user in the handler layer")
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		// logrus.Errorf("Invalid payload for register: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// logrus.Info("Starting to register new user")
	user, err := a.authUsecase.Register(payload)
	if err != nil {
		// logrus.Error("Failed to register user: ", err)
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	// logrus.Info("User has been registered successfully")
	ctx.JSON(http.StatusCreated, user)
}

func (a *AuthController) Route() {
	a.rg.POST("/auth/login", a.loginHandler)
	a.rg.POST("/auth/register", a.registerHandler)
}

func NewAuthController(authUc usecase.AuthUseCase, rg *gin.RouterGroup) *AuthController {
	return &AuthController{authUsecase: authUc, rg: rg}
}

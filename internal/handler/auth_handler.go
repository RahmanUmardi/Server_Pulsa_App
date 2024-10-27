package handler

import (
	"net/http"
	"server-pulsa-app/config"
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

// Login godoc
// @Summary Login user
// @Description Authenticate a user and get JWT token
// @Tags authentication
// @Accept json
// @Produce json
// @Param request body dto.AuthRequest true "Login credentials"
// @Success 200 {object} dto.AuthResponse "Successfully authenticated"
// @Failure 400 {object} dto.ErrorResponse "Invalid input"
// @Failure 401 {object} dto.ErrorResponse "Authentication failed"
// @Router /auth/login [post]
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

// Login godoc
// @Summary Register user
// @Description Create a new user
// @Tags authentication
// @Accept json
// @Produce json
// @Param request body dto.AuthRequest true "Login credentials"
// @Success 201 {object} dto.AuthRegisterRes "Successfully registered"
// @Failure 400 {object} dto.ErrorResponse "Invalid input"
// @Failure 401 {object} dto.ErrorResponse "Authentication failed"
// @Router /auth/register [post]
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
	a.rg.POST(config.Login, a.loginHandler)
	a.rg.POST(config.Register, a.registerHandler)
}

func NewAuthController(authUc usecase.AuthUseCase, rg *gin.RouterGroup, log *logger.Logger) *AuthController {
	return &AuthController{authUsecase: authUc, rg: rg, log: log}
}

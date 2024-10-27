package handler

import (
	"fmt"
	"net/http"
	"server-pulsa-app/config"
	"server-pulsa-app/internal/entity"
	"server-pulsa-app/internal/logger"
	"server-pulsa-app/internal/middleware"
	"server-pulsa-app/internal/usecase"

	// "server-pulsa-app/config"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUc         usecase.UserUsecase
	rg             *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
	log            *logger.Logger
}

func (u *UserHandler) ListHandler(ctx *gin.Context) {
	u.log.Info("Starting to get all user in the handler layer", nil)

	users, err := u.userUc.ListUser()
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	if len(users) > 0 {
		ctx.JSON(http.StatusOK, users)
		return
	}

	reponse := struct {
		Message string
		Data    []entity.User
	}{
		Message: "List of user is empty",
		Data:    users,
	}

	ctx.JSON(http.StatusOK, reponse)
}

func (u *UserHandler) getIdHandler(ctx *gin.Context) {
	u.log.Info("Starting to get user by id in the handler layer", nil)

	id := ctx.Param("id")

	user, err := u.userUc.GetUserByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, fmt.Sprintf("User with id %s not found", id))
		return
	}
	response := struct {
		Message string
		Data    entity.User
	}{
		Message: "Success Get User By Id",
		Data:    user,
	}

	ctx.JSON(http.StatusOK, response)
}

func (u *UserHandler) updateHandler(ctx *gin.Context) {
	u.log.Info("Starting to update user in the handler layer", nil)
	id := ctx.Param("id")
	var payload entity.User
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("User with id %s not found", id))
		return
	}

	payload.Id_user = id

	user, err := u.userUc.UpdateUser(payload)

	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	response := struct {
		Message string
		Data    entity.User
	}{
		Message: "Success Update User",
		Data:    user,
	}

	ctx.JSON(http.StatusOK, response)

}

func (u *UserHandler) deleteHandler(ctx *gin.Context) {
	u.log.Info("Starting to delete user in the handler layer", nil)

	id := ctx.Param("id")
	err := u.userUc.DeleteUser(id)
	if err != nil {
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("User with ID %s not found", id)})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete user", "error": err.Error()})
		return
	}

	response := struct {
		Message string `json:"message"`
	}{
		Message: "User deleted successfully",
	}
	ctx.JSON(http.StatusOK, response)
}

func (u *UserHandler) Route() {
	u.rg.GET(config.GetUserList, u.authMiddleware.RequireToken("admin"), u.ListHandler)
	u.rg.GET(config.GetUser, u.authMiddleware.RequireToken("admin"), u.getIdHandler)
	u.rg.PUT(config.PutUser, u.authMiddleware.RequireToken("admin"), u.updateHandler)
	u.rg.DELETE(config.DeleteUser, u.authMiddleware.RequireToken("admin"), u.deleteHandler)
}

func NewUserHandler(userUc usecase.UserUsecase, authMiddleware middleware.AuthMiddleware, rg *gin.RouterGroup, log *logger.Logger) *UserHandler {
	return &UserHandler{userUc: userUc, authMiddleware: authMiddleware, rg: rg, log: log}
}

package handler

import (
	"fmt"
	"net/http"
	"server-pulsa-app/internal/entity"
	"server-pulsa-app/internal/middleware"
	"server-pulsa-app/internal/usecase"

	"github.com/gin-gonic/gin"
)

// @title User API
// @version 1.0
// @description User management endpoints for the server-pulsa-app
type UserHandler struct {
	userUc         usecase.UserUsecase
	rg             *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
}

// ListUser godoc
// @Summary List all Users
// @Description Get a list of all users
// @Tags products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Success 200 {array} []entity.User "List of users"
// @Failure 401 {object} entity.UserErrorResponse "Unauthorized"
// @Router /users [get]
func (u *UserHandler) ListHandler(ctx *gin.Context) {

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
	id := ctx.Param("id")
	err := u.userUc.DeleteUser(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, fmt.Sprintf("User with id %s not found", id))
		return
	}
	response := struct {
		Message string
	}{
		Message: "Success Delete User",
	}
	ctx.JSON(http.StatusOK, response)
}

func (u *UserHandler) Route() {
	u.rg.GET("/users", u.authMiddleware.RequireToken("admin"), u.ListHandler)
	u.rg.GET("/user/:id", u.authMiddleware.RequireToken("admin"), u.getIdHandler)
	u.rg.PUT("/user/:id", u.authMiddleware.RequireToken("admin"), u.updateHandler)
	u.rg.DELETE("/user/:id", u.authMiddleware.RequireToken("admin"), u.deleteHandler)
}

func NewUserHandler(userUc usecase.UserUsecase, authMiddleware middleware.AuthMiddleware, rg *gin.RouterGroup) *UserHandler {
	return &UserHandler{userUc: userUc, authMiddleware: authMiddleware, rg: rg}
}

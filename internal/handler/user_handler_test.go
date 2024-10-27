package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"server-pulsa-app/internal/entity"
	"server-pulsa-app/internal/logger"
	"server-pulsa-app/internal/mock/middleware_mock"
	"server-pulsa-app/internal/mock/usecase_mock"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type UserHandlerTest struct {
	suite.Suite
	userUc         *usecase_mock.UserUseCaseMock
	router         *gin.Engine
	authMiddleware *middleware_mock.AuthMiddlewareMock
	userHandler    *UserHandler
	log            logger.Logger
}

func (u *UserHandlerTest) SetupTest() {
	u.userUc = new(usecase_mock.UserUseCaseMock)
	u.authMiddleware = new(middleware_mock.AuthMiddlewareMock)

	u.router = gin.Default()
	gin.SetMode(gin.TestMode)

	rg := u.router.Group("/api/v1")

	u.log = logger.NewLogger()
	u.userHandler = NewUserHandler(u.userUc, u.authMiddleware, rg, &u.log)
	u.router.GET("/api/v1/users", u.userHandler.ListHandler)
	u.router.GET("/api/v1/user/:id", u.userHandler.getIdHandler)
	u.router.PUT("/api/v1/user/:id", u.userHandler.updateHandler)
	u.router.DELETE("/api/v1/user/:id", u.userHandler.deleteHandler)
}

func (u *UserHandlerTest) TestUpdate() {
	payload := entity.User{
		Id_user:  "uuid-user-test",
		Username: "testuser",
		Password: "password",
		Role:     "admin",
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		u.T().Fatalf("error '%s' occured when marshaling the payload", err)
	}
	u.userUc.On("UpdateUser", payload).Return(payload, nil)
	request, err := http.NewRequest("PUT", "/api/v1/user/"+payload.Id_user, bytes.NewBuffer(jsonPayload))
	if err != nil {
		u.T().Fatalf("error '%s' occured when creating the request", err)
	}

	w := httptest.NewRecorder()
	u.router.ServeHTTP(w, request)

	u.Equal(http.StatusOK, w.Code)
}

func (u *UserHandlerTest) TestList() {
	u.userUc.On("ListUser").Return([]entity.User{}, nil)

	request, err := http.NewRequest("GET", "/api/v1/users", nil)
	if err != nil {
		u.T().Fatalf("error '%s' occurred when creating the request", err)
	}

	w := httptest.NewRecorder()
	u.router.ServeHTTP(w, request)

	u.Equal(http.StatusOK, w.Code)
}

func (u *UserHandlerTest) TestGet() {
	id := "uuid-user-test"
	u.userUc.On("GetUserByID", id).Return(entity.User{}, nil)
	request, err := http.NewRequest("GET", "/api/v1/user/"+id, nil)
	if err != nil {
		u.T().Fatalf("error '%s' occured when creating the request", err)
	}

	w := httptest.NewRecorder()
	u.router.ServeHTTP(w, request)

	u.Equal(http.StatusOK, w.Code)
}

func (u *UserHandlerTest) TestDelete() {
	id := "uuid-user-test"
	u.userUc.On("DeleteUser", id).Return(nil)
	request, err := http.NewRequest("DELETE", "/api/v1/user/"+id, nil)
	if err != nil {
		u.T().Fatalf("error '%s' occured when creating the request", err)
	}

	w := httptest.NewRecorder()
	u.router.ServeHTTP(w, request)

	u.Equal(http.StatusOK, w.Code)
}

func TestUserHandlerSuite(t *testing.T) {
	suite.Run(t, new(UserHandlerTest))
}

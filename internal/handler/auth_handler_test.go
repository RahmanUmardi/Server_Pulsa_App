package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"server-pulsa-app/internal/entity"
	"server-pulsa-app/internal/entity/dto"
	"server-pulsa-app/internal/logger"
	"server-pulsa-app/internal/mock/usecase_mock"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type AuthHandlerTest struct {
	suite.Suite
	authUc         *usecase_mock.AuthUseCaseMock
	router         *gin.Engine
	AuthController *AuthController
	log            *logger.Logger
}

func (a *AuthHandlerTest) SetupTest() {
	a.authUc = new(usecase_mock.AuthUseCaseMock)

	a.router = gin.Default()
	gin.SetMode(gin.TestMode)

	rg := a.router.Group("/api/v1")

	a.AuthController = NewAuthController(a.authUc, rg, a.log)

	a.AuthController.Route()
}

func (a *AuthHandlerTest) TestLogin() {
	user := entity.User{Username: "testuser", Password: "password"}
	a.authUc.On("Login", user).Return(dto.AuthResponseDto{Token: "some-token"}, nil)

	request, err := http.NewRequest("POST", "/auth/login", bytes.NewBuffer([]byte(`{"username": "testuser", "password": "password"}`)))
	if err != nil {
		a.T().Fail()
	}

	recorder := httptest.NewRecorder()
	a.router.ServeHTTP(recorder, request)

	a.Equal(http.StatusNotFound, recorder.Code)

	var response dto.AuthResponseDto
	_ = json.Unmarshal(recorder.Body.Bytes(), &response)
	a.Equal("", response.Token)
}

func (a *AuthHandlerTest) TestRegister() {
	user := entity.User{Username: "testuser", Password: "password"}
	a.authUc.On("Register", user).Return(user, nil)

	request, err := http.NewRequest("POST", "/auth/register", bytes.NewBuffer([]byte(`{"username": "testuser", "password": "password"}`)))
	if err != nil {
		a.T().Fail()
	}

	recorder := httptest.NewRecorder()
	a.router.ServeHTTP(recorder, request)

	a.Equal(http.StatusNotFound, recorder.Code)

	var response entity.User
	_ = json.Unmarshal(recorder.Body.Bytes(), &response)
	a.Equal("", response.Username)
}

func TestAuthHandlerSuite(t *testing.T) {
	suite.Run(t, new(AuthHandlerTest))
}

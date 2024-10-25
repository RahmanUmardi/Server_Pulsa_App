package handler

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"

	"server-pulsa-app/internal/usecase_mock"
)

type UserHandlerTestSuite struct {
	suite.Suite
	router *gin.Engine
	userUC *usecase_mock.UserUseCaseMock
}

func (u *UserHandlerTestSuite) SetupTest() {
	u.userUC = new(usecase_mock.UserUseCaseMock)

	u.router = gin.Default()
	gin.SetMode(gin.TestMode)

	rg := u.router.Group("/api/v1")

	userUc := NewUserHandler(u.userUC, rg)
	u.router.GET("/users", userUc.ListHandler)
}

func TestUserHandlerSuite(t *testing.T) {
	suite.Run(t, new(UserHandlerTestSuite))
}

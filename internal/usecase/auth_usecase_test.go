package usecase

import (
	"testing"

	"server-pulsa-app/internal/entity"
	"server-pulsa-app/internal/entity/dto"
	"server-pulsa-app/internal/logger"
	"server-pulsa-app/internal/mock/service_mock"
	"server-pulsa-app/internal/mock/usecase_mock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AuthUseCaseTestSuite struct {
	suite.Suite
	authUC          AuthUseCase
	mockUserUsecase *usecase_mock.UserUseCaseMock
	mockJwtService  *service_mock.JwtServiceMock
	log             logger.Logger
}

func (suite *AuthUseCaseTestSuite) SetupTest() {
	suite.mockUserUsecase = new(usecase_mock.UserUseCaseMock)
	suite.mockJwtService = new(service_mock.JwtServiceMock)
	suite.log = logger.NewLogger()
	suite.authUC = NewAuthUseCase(suite.mockUserUsecase, suite.mockJwtService, &suite.log)
}

func (suite *AuthUseCaseTestSuite) TestLogin() {
	user := entity.User{Username: "testuser", Password: "password"}
	suite.mockUserUsecase.On("FindUserByUsernamePassword", "testuser", "password").Return(user, nil)
	suite.mockJwtService.On("CreateToken", user).Return(dto.AuthResponseDto{Token: "mockToken"}, nil)

	response, err := suite.authUC.Login(dto.AuthRequestDto{Username: "testuser", Password: "password"})

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "mockToken", response.Token)

	suite.mockUserUsecase.AssertExpectations(suite.T())
	suite.mockJwtService.AssertExpectations(suite.T())
}

func (suite *AuthUseCaseTestSuite) TestRegister() {
	user := entity.User{Username: "testuser", Password: "password"}
	suite.mockUserUsecase.On("RegisterUser", user).Return(user, nil)

	createdUser, err := suite.authUC.Register(dto.AuthRequestDto{Username: "testuser", Password: "password"})

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), user.Username, createdUser.Username)

	suite.mockUserUsecase.AssertExpectations(suite.T())
}

func TestAuthUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(AuthUseCaseTestSuite))
}

package usecase

import (
	"testing"

	"server-pulsa-app/internal/entity"
	"server-pulsa-app/internal/entity/dto"
	"server-pulsa-app/internal/logger"
	"server-pulsa-app/internal/mock/service_mock"
	"server-pulsa-app/internal/mock/usecase_mock"

	"github.com/stretchr/testify/assert"
)

var log *logger.Logger

func TestAuthUseCase_Login(t *testing.T) {
	mockUserUsecase := new(usecase_mock.UserUseCaseMock)
	mockJwtService := new(service_mock.JwtServiceMock)

	user := entity.User{Username: "testuser", Password: "password"}
	mockUserUsecase.On("FindUserByUsernamePassword", "testuser", "password").Return(user, nil)
	mockJwtService.On("CreateToken", user).Return(dto.AuthResponseDto{Token: "mockToken"}, nil)

	authUC := NewAuthUseCase(mockUserUsecase, mockJwtService, log)

	response, err := authUC.Login(dto.AuthRequestDto{Username: "testuser", Password: "password"})

	assert.NoError(t, err)
	assert.Equal(t, "mockToken", response.Token)

	mockUserUsecase.AssertExpectations(t)
	mockJwtService.AssertExpectations(t)
}

func TestAuthUseCase_Register(t *testing.T) {
	mockUserUsecase := new(usecase_mock.UserUseCaseMock)
	mockJwtService := new(service_mock.JwtServiceMock)

	user := entity.User{Username: "testuser", Password: "password"}
	mockUserUsecase.On("RegisterUser", user).Return(user, nil)

	authUC := NewAuthUseCase(mockUserUsecase, mockJwtService, log)

	createdUser, err := authUC.Register(dto.AuthRequestDto{Username: "testuser", Password: "password"})

	assert.NoError(t, err)
	assert.Equal(t, user.Username, createdUser.Username)

	mockUserUsecase.AssertExpectations(t)
}

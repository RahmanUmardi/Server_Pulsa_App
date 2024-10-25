package usecase

import (
	"testing"

	"server-pulsa-app/internal/entity"
	"server-pulsa-app/internal/entity/dto"
	"server-pulsa-app/internal/mock/service_mock"
	"server-pulsa-app/internal/mock/usecase_mock"

	"github.com/stretchr/testify/assert"
)

func TestAuthUseCase_Login(t *testing.T) {
	mockUserUsecase := new(usecase_mock.UserUseCaseMock)
	mockJwtService := new(service_mock.JwtServiceMock)

	// Setup
	user := entity.User{Username: "testuser", Password: "password"}
	mockUserUsecase.On("FindUserByUsernamePassword", "testuser", "password").Return(user, nil)
	mockJwtService.On("CreateToken", user).Return(dto.AuthResponseDto{Token: "mockToken"}, nil)

	authUC := NewAuthUseCase(mockUserUsecase, mockJwtService)

	// Act
	response, err := authUC.Login(dto.AuthRequestDto{Username: "testuser", Password: "password"})

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, "mockToken", response.Token)

	// Verify
	mockUserUsecase.AssertExpectations(t)
	mockJwtService.AssertExpectations(t)
}

func TestAuthUseCase_Register(t *testing.T) {
	mockUserUsecase := new(usecase_mock.UserUseCaseMock)
	mockJwtService := new(service_mock.JwtServiceMock)

	// Setup
	user := entity.User{Username: "testuser", Password: "password"}
	mockUserUsecase.On("RegisterUser", user).Return(user, nil)

	authUC := NewAuthUseCase(mockUserUsecase, mockJwtService)

	// Act
	createdUser, err := authUC.Register(dto.AuthRequestDto{Username: "testuser", Password: "password"})

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, user.Username, createdUser.Username)

	// Verify
	mockUserUsecase.AssertExpectations(t)
}

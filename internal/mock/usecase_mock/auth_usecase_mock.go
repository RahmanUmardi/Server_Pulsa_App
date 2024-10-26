package usecase_mock

import (
	"server-pulsa-app/internal/entity"
	"server-pulsa-app/internal/entity/dto"

	"github.com/stretchr/testify/mock"
)

type AuthUseCaseMock struct {
	mock.Mock
}

func (a *AuthUseCaseMock) Login(payload dto.AuthRequestDto) (dto.AuthResponseDto, error) {
	args := a.Called(payload)
	return args.Get(0).(dto.AuthResponseDto), args.Error(1)
}

func (a *AuthUseCaseMock) Register(payload dto.AuthRequestDto) (entity.User, error) {
	args := a.Called(payload)
	return args.Get(0).(entity.User), args.Error(1)
}

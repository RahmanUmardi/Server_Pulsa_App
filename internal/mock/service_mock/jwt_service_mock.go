package service_mock

import (
	"server-pulsa-app/internal/entity"
	"server-pulsa-app/internal/entity/dto"
	"server-pulsa-app/internal/shared/model"

	"github.com/stretchr/testify/mock"
)

type JwtServiceMock struct {
	mock.Mock
}

func (j *JwtServiceMock) CreateToken(user entity.User) (dto.AuthResponseDto, error) {
	args := j.Called(user)
	return args.Get(0).(dto.AuthResponseDto), args.Error(1)
}

func (j *JwtServiceMock) ValidateToken(tokenString string) (*model.Claim, error) {
	args := j.Called(tokenString)
	return args.Get(0).(*model.Claim), args.Error(1)
}

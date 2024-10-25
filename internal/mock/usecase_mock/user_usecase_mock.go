package usecase_mock

import (
	"server-pulsa-app/internal/entity"

	"github.com/stretchr/testify/mock"
)

type UserUseCaseMock struct {
	mock.Mock
}

func (u *UserUseCaseMock) RegisterUser(payload entity.User) (entity.User, error) {
	args := u.Called(payload)
	return args.Get(0).(entity.User), args.Error(1)
}

func (u *UserUseCaseMock) GetUserByUsername(username string) (entity.User, error) {
	args := u.Called(username)
	return args.Get(0).(entity.User), args.Error(1)
}

func (u *UserUseCaseMock) GetUserByID(id string) (entity.User, error) {
	args := u.Called(id)
	return args.Get(0).(entity.User), args.Error(1)
}

func (u *UserUseCaseMock) FindUserByUsernamePassword(username, password string) (entity.User, error) {
	args := u.Called(username, password)
	return args.Get(0).(entity.User), args.Error(1)
}

func (u *UserUseCaseMock) ListUser() ([]entity.User, error) {
	args := u.Called()
	return args.Get(0).([]entity.User), args.Error(1)
}

func (u *UserUseCaseMock) UpdateUser(payload entity.User) (entity.User, error) {
	args := u.Called(payload)
	return args.Get(0).(entity.User), args.Error(1)
}

func (u *UserUseCaseMock) DeleteUser(id string) error {
	args := u.Called(id)
	return args.Error(0)
}

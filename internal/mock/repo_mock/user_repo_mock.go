package repo_mock

import (
	"server-pulsa-app/internal/entity"

	"github.com/stretchr/testify/mock"
)

type UserRepoMock struct {
	mock.Mock
}

func (u *UserRepoMock) CreateUser(payload entity.User) (entity.User, error) {
	args := u.Called(payload)
	return args.Get(0).(entity.User), args.Error(1)
}

func (u *UserRepoMock) GetUserByUsername(username string) (entity.User, error) {
	args := u.Called(username)
	return args.Get(0).(entity.User), args.Error(1)
}

func (u *UserRepoMock) GetUserByID(id string) (entity.User, error) {
	args := u.Called(id)
	return args.Get(0).(entity.User), args.Error(1)
}

func (u *UserRepoMock) ListUser() ([]entity.User, error) {
	args := u.Called()
	return args.Get(0).([]entity.User), args.Error(1)
}

func (u *UserRepoMock) UpdateUser(payload entity.User) (entity.User, error) {
	args := u.Called(payload)
	return args.Get(0).(entity.User), args.Error(1)
}

func (u *UserRepoMock) DeleteUser(id string) error {
	args := u.Called(id)
	return args.Error(0)
}

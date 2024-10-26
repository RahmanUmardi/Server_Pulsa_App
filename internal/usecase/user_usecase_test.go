package usecase

import (
	"server-pulsa-app/internal/entity"
	"server-pulsa-app/internal/logger"
	"server-pulsa-app/internal/mock/repo_mock"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
)

type userUsecaseTestSuite struct {
	suite.Suite
	mockUserRepository *repo_mock.UserRepoMock
	UserUseCase        UserUsecase
	log                logger.Logger
}

func (u *userUsecaseTestSuite) SetupTest() {
	u.mockUserRepository = new(repo_mock.UserRepoMock)
	u.log = logger.NewLogger()
	u.UserUseCase = NewUserUsecase(u.mockUserRepository, &u.log)
}

func (u *userUsecaseTestSuite) TestRegisterUser_Success() {
	username := "Test User"
	user := entity.User{
		Id_user:  "1",
		Username: username,
		Password: "Test Password",
		Role:     "Test Role",
	}

	u.mockUserRepository.On("GetUserByUsername", username).Return(entity.User{}, nil).Once()

	u.mockUserRepository.On("CreateUser", mock.Anything).Return(user, nil).Once()

	user, err := u.UserUseCase.RegisterUser(user)

	u.NoError(err)
	u.Equal("1", user.Id_user)
}

func (u *userUsecaseTestSuite) TestListAll_Success() {
	user := []entity.User{
		{
			Id_user:  "1",
			Username: "Test User",
			Password: "Test Password",
			Role:     "Test Role",
		},
		{
			Id_user:  "2",
			Username: "Test User",
			Password: "Test Password",
			Role:     "Test Role",
		},
	}

	u.mockUserRepository.On("ListUser").Return(user, nil).Once()

	userList, err := u.UserUseCase.ListUser()

	u.Nil(err)
	u.Equal(user, userList)
}

func (u *userUsecaseTestSuite) TestGetUserById_Success() {
	id := "1"

	user := entity.User{
		Id_user:  "1",
		Username: "Test User",
		Password: "Test Password",
		Role:     "Test Role",
	}

	u.mockUserRepository.On("GetUserByID", id).Return(user, nil).Once()

	userFound, err := u.UserUseCase.GetUserByID(id)

	u.Nil(err)
	u.Equal(user, userFound)
}

func (u *userUsecaseTestSuite) TestUpdateUser_Success() {
	id := "1"
	updatedUser := entity.User{
		Id_user:  "1",
		Username: "Test User",
		Password: hashPassword("Test Password"),
		Role:     "Test Role",
	}

	u.mockUserRepository.On("GetUserByID", id).Return(entity.User{
		Id_user:  id,
		Username: "Test User",
		Password: hashPassword("test_password"),
		Role:     "Test Role",
	}, nil).Once()

	u.mockUserRepository.On("UpdateUser", mock.Anything).Return(updatedUser, nil).Once()

	userUpdated, err := u.UserUseCase.UpdateUser(updatedUser)

	u.Nil(err)
	u.Equal(updatedUser.Id_user, userUpdated.Id_user)
}

func (u *userUsecaseTestSuite) TestDeleteUser_Success() {
	id := "1"

	u.mockUserRepository.On("GetUserByID", id).Return(entity.User{
		Id_user:  id,
		Username: "Test User",
		Password: hashPassword("Test Password"),
		Role:     "Test Role",
	}, nil).Once()

	u.mockUserRepository.On("DeleteUser", id).Return(nil).Once()

	err := u.UserUseCase.DeleteUser(id)

	u.Nil(err)
}

func hashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hashedPassword)
}

func TestUserUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(userUsecaseTestSuite))
}

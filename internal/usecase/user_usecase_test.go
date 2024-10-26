package usecase

import (
	"server-pulsa-app/internal/entity"
	"server-pulsa-app/internal/mock/repo_mock"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

func TestRegisterUser(t *testing.T) {
	userRepo := new(repo_mock.UserRepoMock)
	useCase := NewUserUsecase(userRepo)

	username := "test"
	user := entity.User{
		Id_user:  "1adc",
		Username: username,
		Password: "password",
		Role:     "employee",
	}

	userRepo.On("GetUserByUsername", username).Return(entity.User{}, nil)

	userRepo.On("CreateUser", mock.Anything).Return(user, nil)

	result, err := useCase.RegisterUser(user)
	assert.NoError(t, err)
	assert.Equal(t, "1adc", result.Id_user)

	userRepo.AssertExpectations(t)
}

func TestListUser(t *testing.T) {
	mockRepo := new(repo_mock.UserRepoMock)
	useCase := NewUserUsecase(mockRepo)

	user := []entity.User{
		{
			Id_user:  "1",
			Username: "test",
			Password: "test",
			Role:     "test",
		},

		{
			Id_user:  "2",
			Username: "test2",
			Password: "test2",
			Role:     "test2",
		},
	}

	mockRepo.On("ListUser").Return(user, nil)

	result, err := useCase.ListUser()
	assert.NoError(t, err)
	assert.Len(t, result, len(user))

	mockRepo.AssertExpectations(t)
}

func TestGetUserById(t *testing.T) {
	mockRepo := new(repo_mock.UserRepoMock)
	useCase := NewUserUsecase(mockRepo)

	user := entity.User{
		Id_user:  "1",
		Username: "test",
		Password: "test",
		Role:     "test",
	}

	mockRepo.On("GetUserByID", "1").Return(user, nil)

	result, err := useCase.GetUserByID("1")
	assert.NoError(t, err)
	assert.Equal(t, user.Id_user, result.Id_user)

	mockRepo.AssertExpectations(t)
}

func TestFindUserByUsernamePassword(t *testing.T) {
	mockRepo := new(repo_mock.UserRepoMock)
	useCase := NewUserUsecase(mockRepo)

	username := "test"
	password := "correct_password"
	hashedPassword := hashPassword(entity.User{}, password)

	user := entity.User{
		Id_user:  "1",
		Username: username,
		Password: hashedPassword,
		Role:     "user",
	}

	mockRepo.On("GetUserByUsername", username).Return(user, nil)

	foundUser, err := useCase.FindUserByUsernamePassword(username, password)
	assert.NoError(t, err)
	assert.Equal(t, user.Id_user, foundUser.Id_user)

	mockRepo.AssertExpectations(t)
}

func hashPassword(user entity.User, password string) string {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hashedPassword)
}

func TestGetUserByUsername(t *testing.T) {
	mockRepo := new(repo_mock.UserRepoMock)
	useCase := NewUserUsecase(mockRepo)

	user := entity.User{
		Id_user:  "1",
		Username: "test",
		Password: "test",
		Role:     "test",
	}

	mockRepo.On("GetUserByUsername", "test").Return(user, nil)

	result, err := useCase.GetUserByUsername("test")
	assert.NoError(t, err)
	assert.Equal(t, user.Id_user, result.Id_user)

	mockRepo.AssertExpectations(t)
}

func TestUpdateUser(t *testing.T) {
	mockRepo := new(repo_mock.UserRepoMock)
	useCase := NewUserUsecase(mockRepo)

	userId := "1"
	updatedUser := entity.User{
		Id_user:  userId,
		Username: "updated_username",
		Password: hashPassword(entity.User{}, "correct_password"),
		Role:     "user",
	}

	mockRepo.On("GetUserByID", userId).Return(entity.User{
		Id_user:  userId,
		Username: "test_username",
		Password: hashPassword(entity.User{}, "correct_password"),
		Role:     "user",
	}, nil)

	mockRepo.On("UpdateUser", updatedUser).Return(updatedUser, nil)

	returnedUser, err := useCase.UpdateUser(updatedUser)
	assert.NoError(t, err)
	assert.Equal(t, updatedUser.Id_user, returnedUser.Id_user)

	mockRepo.AssertExpectations(t)
}

func TestDeleteUser(t *testing.T) {
	mockRepo := new(repo_mock.UserRepoMock)
	useCase := NewUserUsecase(mockRepo)

	userId := "1"

	mockRepo.On("GetUserByID", userId).Return(entity.User{
		Id_user:  userId,
		Username: "test_username",
		Password: hashPassword(entity.User{}, "correct_password"),
		Role:     "user",
	}, nil)

	mockRepo.On("DeleteUser", userId).Return(nil)

	err := useCase.DeleteUser(userId)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

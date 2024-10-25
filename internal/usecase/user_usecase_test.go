package usecase

import (
	"fmt"
	"server-pulsa-app/internal/entity"
	"server-pulsa-app/internal/mock/repo_mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterUser(t *testing.T) {
	userRepo := new(repo_mock.UserRepoMock)
	useCase := NewUserUsecase(userRepo)

	userRepo.On("GetUserByUsername", "test").Return(entity.User{Username: "test", Password: "password"}, nil)

	user := entity.User{
		Id_user:  "1adc",
		Username: "test",
		Password: "test",
		Role:     "test",
	}

	userRepo.On("CreateUser", user).Return(user, fmt.Errorf("Username %s already exists", user.Username))

	result, err := useCase.RegisterUser(user)
	assert.Error(t, err)
	assert.Equal(t, user.Id_user, result.Id_user)

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

	mockRepo.On("GetUserById", "1").Return(user, nil)

	result, err := useCase.GetUserByID("1")
	assert.NoError(t, err)
	assert.Equal(t, user.Id_user, result.Id_user)

	mockRepo.AssertExpectations(t)
}

func TestFindUserByUsernamePassword(t *testing.T) {
	mockRepo := new(repo_mock.UserRepoMock)
	useCase := NewUserUsecase(mockRepo)

	user := entity.User{
		Id_user:  "1",
		Username: "test",
		Password: "test",
		Role:     "test",
	}

	mockRepo.On("FindUserByUsernamePassword", "test", "test").Return(user, nil)

	result, err := useCase.FindUserByUsernamePassword("test", "test")
	assert.NoError(t, err)
	assert.Equal(t, user.Id_user, result.Id_user)

	mockRepo.AssertExpectations(t)
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

	user := entity.User{
		Id_user:  "1",
		Username: "test",
		Password: "test",
		Role:     "test",
	}

	mockRepo.On("UpdateUser", user, user).Return(user, nil)

	result, err := useCase.UpdateUser(user)
	assert.NoError(t, err)
	assert.Equal(t, user.Id_user, result.Id_user)

	mockRepo.AssertExpectations(t)
}

func TestDeleteUser(t *testing.T) {
	mockRepo := new(repo_mock.UserRepoMock)
	useCase := NewUserUsecase(mockRepo)

	user := entity.User{
		Id_user:  "1",
		Username: "test",
		Password: "test",
		Role:     "test",
	}

	mockRepo.On("DeleteUser", user).Return(user, nil)

	err := useCase.DeleteUser("1")
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

// type UserUseCaseTestSuite struct {
// 	suite.Suite
// 	mockUserRepo *repo_mock.UserRepoMock
// 	userUC      UserUsecase
// }

// func (u *UserUseCaseTestSuite) SetupTest() {
// 	u.mockUserRepo = new(repo_mock.UserRepoMock)
//     u.UserUsecase = NewUserUsecase(u.mockUserRepo)
// }

// func (u *UserUseCaseTestSuite) TestRegisterUser_Failed(t *testing.T) {
// 	newUser := entity.User{
// 		Id_user:  "1",
// 		Username: "test",
// 		Password: "test",
// 		Role:     "test",
// 	}

// 	u.userUseCase.On("RegisterUser", newUser).Return(entity.User{
// 		Id_user:  "1",
// 		Username: "test",
// 		Password: "test",
// 		Role:     "test",
// 	}, nil).Once()

// 	_, err := u.userUseCase.RegisterUser(newUser)

// 	u.userUseCase.AssertExpectations(u.T())

// 	assert.Nil(u.T(), err)
// }

// func (u *UserUseCaseTestSuite) TestRegisterUser_Success(t *testing.T) {
// 	newUser := entity.User{
// 		Id_user:  "1",
// 		Username: "test",
// 		Password: "test",
// 		Role:     "test",
// 	}

// 	u.userUseCase.On("RegisterUser", newUser).Return(entity.User{
// 		Id_user:  "1",
// 		Username: "test",
// 		Password: "test",
// 		Role:     "test",
// 	}, nil).Once()

// 	_, err := u.userUseCase.RegisterUser(newUser)

// 	u.userUseCase.AssertExpectations(u.T())

// 	assert.Nil(u.T(), err)
// }

// func (u *UserUseCaseTestSuite) TestListUser_Failed(t *testing.T) {

// 	user := []entity.User{
// 		{
// 			Id_user:  "1",
// 			Username: "test",
// 			Password: "test",
// 			Role:     "test",
// 		},
// 		{
// 			Id_user:  "2",
// 			Username: "tost",
// 			Password: "tost",
// 			Role:     "tost",
// 		},
// 	}

// 	u.userUseCase.On("ListUser").Return(user, fmt.Errorf("failed")).Once()

// 	_, err := u.userUseCase.ListUser()

// 	u.userUseCase.AssertExpectations(u.T())

// 	assert.NotNil(u.T(), err)

// }

// func (u *UserUseCaseTestSuite) TestListUser_Success(t *testing.T) {

// 	user := []entity.User{
// 		{
// 			Id_user:  "1",
// 			Username: "test",
// 			Password: "test",
// 			Role:     "tes",
// 		},
// 		{
// 			Id_user:  "2",
// 			Username: "tost",
// 			Password: "tost",
// 			Role:     "tost",
// 		},
// 	}

// 	u.userUseCase.On("ListUser").Return(user, nil).Once()

// 	_, err := u.userUseCase.ListUser()

// 	u.userUseCase.AssertExpectations(u.T())

// 	assert.Nil(u.T(), err)

// }

// func (u *UserUseCaseTestSuite) TestFindUserByUsernamePassword_Failed(t *testing.T) {

// 	u.userUseCase.On("FindUserByUsernamePassword", "test", "test").Return(entity.User{
// 		Id_user:  "1",
// 		Username: "test",
// 		Password: "test",
// 		Role:     "test",
// 	}).Once()

// 	_, err := u.userUseCase.FindUserByUsernamePassword("test", "test")

// 	u.userUseCase.AssertExpectations(u.T())

// 	assert.NotNil(u.T(), err)

// }

// func (u *UserUseCaseTestSuite) TestFindUserByUsernamePassword_Success(t *testing.T) {

// 	u.userUseCase.On("FindUserByUsernamePassword", "test", "test").Return(entity.User{
// 		Id_user:  "1",
// 		Username: "test",
// 		Password: "test",
// 		Role:     "test",
// 	}).Once()

// 	_, err := u.userUseCase.FindUserByUsernamePassword("test", "test")

// 	u.userUseCase.AssertExpectations(u.T())

// 	assert.Nil(u.T(), err)

// }

// func (u *UserUseCaseTestSuite) TestGetUserById_Success(t *testing.T) {

// 	u.userUseCase.On("GetUserbyId", "1").Return(entity.User{
// 		Id_user:  "1",
// 		Username: "test",
// 		Password: "test",
// 		Role:     "test",
// 	}, nil).Once()

// 	_, err := u.userUseCase.GetUserByID("1")

// 	u.userUseCase.AssertExpectations(u.T())

// 	assert.Nil(u.T(), err)

// }

// func (u *UserUseCaseTestSuite) TestGetUserById_Failed(t *testing.T) {

// 	u.userUseCase.On("GetUserbyId", "1").Return(entity.User{
// 		Id_user:  "1",
// 		Username: "test",
// 		Password: "test",
// 		Role:     "test",
// 	}, nil).Once()

// 	_, err := u.userUseCase.GetUserByID("1")

// 	u.userUseCase.AssertExpectations(u.T())

// 	assert.NotNil(u.T(), err)

// }

// func (u *UserUseCaseTestSuite) TestUpdateUser_Success(t *testing.T) {

// 	newUser := entity.User{
// 		Id_user:  "1",
// 		Username: "test",
// 		Password: "test",
// 		Role:     "test",
// 	}

// 	u.userUseCase.On("UpdateUser", newUser).Return(entity.User{
// 		Id_user:  "1",
// 		Username: "test",
// 		Password: "test",
// 		Role:     "test",
// 	}, nil).Once()

// 	_, err := u.userUseCase.UpdateUser(newUser)

// 	u.userUseCase.AssertExpectations(u.T())

// 	assert.Nil(u.T(), err)

// }

// func (u *UserUseCaseTestSuite) TestUpdateUser_Failed(t *testing.T) {

// 	newUser := entity.User{
// 		Id_user:  "1",
// 		Username: "test",
// 		Password: "test",
// 		Role:     "test",
// 	}

// 	u.userUseCase.On("UpdateUser", newUser).Return(entity.User{
// 		Id_user:  "1",
// 		Username: "test",
// 		Password: "test",
// 		Role:     "test",
// 	}, fmt.Errorf("failed")).Once()

// 	_, err := u.userUseCase.UpdateUser(newUser)

// 	u.userUseCase.AssertExpectations(u.T())

// 	assert.NotNil(u.T(), err)

// }

// func (u *UserUseCaseTestSuite) TestDeleteUser_Success(t *testing.T) {

// 	u.userUseCase.On("DeleteUser", "1").Return(nil).Once()

// 	err := u.userUseCase.DeleteUser("1")

// 	u.userUseCase.AssertExpectations(u.T())

// 	assert.Nil(u.T(), err)

// }

// func (u *UserUseCaseTestSuite) TestDeleteUser_Failed(t *testing.T) {

// 	u.userUseCase.On("DeleteUser", "1").Return(fmt.Errorf("failed")).Once()

// 	err := u.userUseCase.DeleteUser("1")

// 	u.userUseCase.AssertExpectations(u.T())

// 	assert.NotNil(u.T(), err)

// 	// u.userUseCase.On("GetByID", "1").Return(entity.User{}, nil)
// 	// u.userUseCase.On("Delete", "1").Return(fmt.Errorf("failed")).Once()

// 	// err = u.userUseCase.DeleteUser("1")
// 	// assert.NoError(t, err)
// 	// assert.Equal(t, , product.ID)

// 	// u.userUseCase.AssertExpectations(t)

// }

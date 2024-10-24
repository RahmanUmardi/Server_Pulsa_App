package usecase

import (
	"fmt"
	"server-pulsa-app/internal/entity"
	"server-pulsa-app/internal/logger"
	"server-pulsa-app/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

var logUser = logger.GetLogger()

type UserUsecase interface {
	RegisterUser(user entity.User) (entity.User, error)
	GetUserByID(id string) (entity.User, error)
	FindUserByUsernamePassword(username, password string) (entity.User, error)
	UpdateUser(user entity.User) (entity.User, error)
	DeleteUser(id string) error
}

type userUsecase struct {
	UserRepository repository.UserRepository
}

func (u *userUsecase) RegisterUser(user entity.User) (entity.User, error) {
	logUser.Info("Starting to create a new user in the usecase layer")

	existUser, _ := u.UserRepository.GetUserByUsername(user.Username)
	logUser.Info("Starting to validate a new user")
	if existUser.Username == user.Username {
		logUser.Error("Username already exist")
		return entity.User{}, fmt.Errorf("username already exist")
	}

	logUser.Info("Starting to set default role for new user")
	user.Role = "employee"
	logUser.Info("Starting to hash the password")
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		logUser.Error("Failed to hash password: ", err)
		return entity.User{}, err
	}

	user.Password = string(hash)

	logUser.Info("Starting to create a new user in the repository layer")
	return u.UserRepository.CreateUser(user)
}

func (u *userUsecase) GetUserByUsername(username string) (entity.User, error) {
	logUser.Info("Starting to retrieve a user by username in the usecase layer")
	return u.UserRepository.GetUserByUsername(username)
}

func (u *userUsecase) GetUserByID(id string) (entity.User, error) {
	logUser.Info("Starting to retrieve a user by id in the usecase layer")
	return u.UserRepository.GetUserByID(id)
}

func (u *userUsecase) FindUserByUsernamePassword(username, password string) (entity.User, error) {
	logUser.Info("Starting to authenticate a user in the usecase layer")

	userExist, err := u.UserRepository.GetUserByUsername(username)
	if err != nil {
		logUser.Errorf("User ID %s not found: %v", userExist.Id_user, err)
		return entity.User{}, fmt.Errorf("user doesn't exists")
	}

	logUser.Info("Starting to validate password")
	err = bcrypt.CompareHashAndPassword([]byte(userExist.Password), []byte(password))
	if err != nil {
		logUser.Error("Password doesn't match")
		return entity.User{}, fmt.Errorf("password doesn't match")
	}

	logUser.Infof("User ID %s has been authenticated successfully: ", userExist.Id_user)
	return userExist, nil
}

func (u *userUsecase) UpdateUser(user entity.User) (entity.User, error) {
	logUser.Info("Starting to update a user in the usecase layer")

	_, err := u.UserRepository.UpdateUser(user)
	if err != nil {
		logUser.Error("Failed to update user: ", err)
		return entity.User{}, fmt.Errorf("failed to update user: %v", err)
	}

	logUser.Infof("User ID %s has been updated successfully: ", user.Id_user)
	return user, nil
}

func (u *userUsecase) DeleteUser(id string) error {
	logUser.Info("Starting to delete a user in the usecase layer")

	err := u.UserRepository.DeleteUser(id)
	if err != nil {
		logUser.Error("Failed to delete user: ", err)
		return fmt.Errorf("failed to delete user: %v", err)
	}

	logUser.Infof("User ID %s has been deleted successfully: ", id)
	return nil
}

func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{UserRepository: userRepository}
}

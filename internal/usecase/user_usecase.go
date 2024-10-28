package usecase

import (
	"fmt"
	"server-pulsa-app/internal/entity"
	"server-pulsa-app/internal/logger"
	"server-pulsa-app/internal/repository"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	RegisterUser(user entity.User) (entity.User, error)
	GetUserByID(id string) (entity.User, error)
	ListUser() ([]entity.User, error)
	GetUserByUsername(username string) (entity.User, error)
	FindUserByUsernamePassword(username, password string) (entity.User, error)
	UpdateUser(payload entity.User) (entity.User, error)
	DeleteUser(id string) error
}

type userUsecase struct {
	UserRepository repository.UserRepository
	log            *logger.Logger
}

func (u *userUsecase) RegisterUser(user entity.User) (entity.User, error) {
	u.log.Info("Starting to create a new user in the usecase layer", nil)

	existUser, _ := u.UserRepository.GetUserByUsername(user.Username)
	u.log.Info("Starting to validate a new user", nil)
	if existUser.Username == user.Username {
		u.log.Error("Username already exist", existUser.Username)
		return entity.User{}, fmt.Errorf("username already exist")
	}

	u.log.Info("Starting to set default role for new user", nil)
	user.Role = "employee"
	u.log.Info("Starting to hash the password", nil)
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		u.log.Error("Failed to hash password: ", err)
		return entity.User{}, err
	}

	user.Password = string(hash)

	u.log.Info("Starting to create a new user in the repository layer", nil)
	return u.UserRepository.CreateUser(user)
}

func (u *userUsecase) GetUserByUsername(username string) (entity.User, error) {
	u.log.Info("Starting to retrieve a user by username in the usecase layer", nil)
	return u.UserRepository.GetUserByUsername(username)
}

func (u *userUsecase) ListUser() ([]entity.User, error) {
	logrus.Info("Starting to get list user in the usecase layer")
	return u.UserRepository.ListUser()
}

func (u *userUsecase) GetUserByID(id string) (entity.User, error) {
	u.log.Info("Starting to retrieve a user by id in the usecase layer", nil)
	return u.UserRepository.GetUserByID(id)
}

func (u *userUsecase) FindUserByUsernamePassword(username, password string) (entity.User, error) {
	u.log.Info("Starting to authenticate a user in the usecase layer", nil)

	userExist, err := u.UserRepository.GetUserByUsername(username)
	if err != nil {
		u.log.Error("User ID %s not found: %v", userExist.Id_user)
		return entity.User{}, fmt.Errorf("user doesn't exists")
	}

	u.log.Info("Starting to validate password", nil)
	err = bcrypt.CompareHashAndPassword([]byte(userExist.Password), []byte(password))
	if err != nil {
		u.log.Error("Password doesn't match", err)
		return entity.User{}, fmt.Errorf("password doesn't match")
	}

	u.log.Info("User ID %s has been authenticated successfully: ", userExist.Id_user)
	return userExist, nil
}

func (u *userUsecase) UpdateUser(user entity.User) (entity.User, error) {
	u.log.Info("Starting to update a user in the usecase layer", nil)

	_, err := u.UserRepository.GetUserByID(user.Id_user)
	if err != nil {
		u.log.Error("User ID %s not found: %v", user.Id_user)
		return entity.User{}, fmt.Errorf("user ID %s not found", user.Id_user)
	}
	u.log.Info("Starting to hash the password", nil)
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		u.log.Error("Failed to hash password: ", err)
		return entity.User{}, fmt.Errorf("failed to hash password: %v", err)
	}
	user.Password = string(hash)

	updatedUser, err := u.UserRepository.UpdateUser(user)
	if err != nil {
		u.log.Error("Failed to update user: ", err)
		return entity.User{}, fmt.Errorf("failed to update user: %v", err)
	}

	u.log.Info("User ID %s has been updated successfully: ", user.Id_user)
	return updatedUser, nil
}

func (u *userUsecase) DeleteUser(id string) error {
	u.log.Info("Starting to delete a user in the usecase layer", nil)

	_, err := u.UserRepository.GetUserByID(id)
	if err != nil {
		u.log.Error("User ID %s not found: %v", id)
		return fmt.Errorf("user ID %s not found", id)
	}

	err = u.UserRepository.DeleteUser(id)
	if err != nil {
		u.log.Error("Failed to delete user: ", err)
		return fmt.Errorf("failed to delete user: %v", err)
	}

	u.log.Info("User ID %s has been deleted successfully: ", id)
	return nil
}

func NewUserUsecase(userRepository repository.UserRepository, log *logger.Logger) UserUsecase {
	return &userUsecase{UserRepository: userRepository, log: log}
}

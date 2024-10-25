package usecase

import (
	"fmt"
	"server-pulsa-app/internal/entity"
	"server-pulsa-app/internal/logger"
	"server-pulsa-app/internal/repository"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

var logUser = logger.GetLogger()

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
}

func (u *userUsecase) RegisterUser(user entity.User) (entity.User, error) {
	logrus.Info("Starting to register user in the usecase layer")
	existUser, _ := u.UserRepository.GetUserByUsername(user.Username)
	if existUser.Username == user.Username {
		return entity.User{}, fmt.Errorf("username already exists")
	}
	user.Role = "employee"
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return entity.User{}, err
	}
	user.Password = string(hash)

	return u.UserRepository.CreateUser(user)
}

func (u *userUsecase) GetUserByUsername(username string) (entity.User, error) {
	logrus.Info("Starting to get user by username in the usecaselayer")
	return u.UserRepository.GetUserByUsername(username)
}

func (u *userUsecase) ListUser() ([]entity.User, error) {
	logrus.Info("Starting to get list user in the usecase layer")
	return u.UserRepository.ListUser()
}

func (u *userUsecase) GetUserByID(id string) (entity.User, error) {
	return u.UserRepository.GetUserByID(id)
}

func (u *userUsecase) FindUserByUsernamePassword(username, password string) (entity.User, error) {
	logrus.Info("Starting find user by username password in the usecase")
	userExist, err := u.UserRepository.GetUserByUsername(username)
	if err != nil {
		return entity.User{}, fmt.Errorf("user doesn't exists")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userExist.Password), []byte(password))
	if err != nil {
		return entity.User{}, fmt.Errorf("password doesn't match")
	}

	return userExist, nil
}

func (u *userUsecase) UpdateUser(payload entity.User) (entity.User, error) {
	logrus.Info("Starting update user in the usecase layer")
	user, err := u.UserRepository.GetUserByID(payload.Id_user)
	if err != nil {
		return entity.User{}, fmt.Errorf("user ID of \\'%s\\' not found", payload.Id_user)
	}
	_, err = u.UserRepository.UpdateUser(user, payload)
	if err != nil {
		return entity.User{}, fmt.Errorf("user ID of \\'%s\\' not updated", payload.Id_user)
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return entity.User{}, err
	}
	user.Password = string(hash)
	return u.UserRepository.UpdateUser(user, payload)
}

func (u *userUsecase) DeleteUser(id string) error {
	logrus.Info("Starting delete user in the usecase layer")
	_, err := u.UserRepository.GetUserByID(id)
	if err != nil {
		return fmt.Errorf("merchant ID of \\%s\\ not found", err)
	}
	return u.UserRepository.DeleteUser(id)
}

func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{UserRepository: userRepository}
}

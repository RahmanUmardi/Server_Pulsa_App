package usecase

import (
	"fmt"
	"server-pulsa-app/internal/entity"
	"server-pulsa-app/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	RegisterUser(user entity.User) (entity.User, error)
	GetUserByID(id string) (entity.User, error)
	UpdateUser(user entity.User) (entity.User, error)
	DeleteUser(id string) error
}

type userUsecase struct {
	UserRepository repository.UserRepository
}

func (u *userUsecase) RegisterUser(user entity.User) (entity.User, error) {
	exitUser, err := u.UserRepository.GetUserByUsername(user.Username)
	if exitUser.Username != user.Username {
		return entity.User{}, fmt.Errorf("username already exist")
	}
	user.Role = "user"
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return entity.User{}, err
	}
	user.Password = string(hash)

	return u.UserRepository.RegisterUser(user)
}

func (u *userUsecase) GetUserByUsername(username string) (entity.User, error) {
	return u.UserRepository.GetUserByUsername(username)
}

func (u *userUsecase) GetUserByID(id string) (entity.User, error) {
	return u.UserRepository.GetUserByID(id)
}

func (u *userUsecase) UpdateUser(user entity.User) (entity.User, error) {
	_, err := u.UserRepository.UpdateUser(user)
	if err != nil {
		return entity.User{}, fmt.Errorf("failed to update user: %v", err)
	}
	return user, nil
}

func (u *userUsecase) DeleteUser(id string) error {
	err := u.UserRepository.DeleteUser(id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
	}
	return nil
}

func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{UserRepository: userRepository}
}

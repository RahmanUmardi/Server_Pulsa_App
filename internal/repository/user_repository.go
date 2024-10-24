package repository

import (
	"database/sql"
	"server-pulsa-app/internal/entity"
	"server-pulsa-app/internal/logger"
)

type UserRepository interface {
	CreateUser(user entity.User) (entity.User, error)
	GetUserByID(id string) (entity.User, error)
	GetUserByUsername(username string) (entity.User, error)
	UpdateUser(user entity.User) (entity.User, error)
	DeleteUser(id string) error
}

var logUser = logger.GetLogger()

type userRepository struct {
	db *sql.DB
}

func (u *userRepository) CreateUser(user entity.User) (entity.User, error) {
	logUser.Info("Starting to create a new user in the repository layer")

	err := u.db.QueryRow(`INSERT INTO mst_user (username, password, role) VALUES ($1, $2, $3) RETURNING id_user`, user.Username, user.Password, user.Role).Scan(&user.Id_user)

	if err != nil {
		logUser.Error("Failed to create the user: ", err)
		return entity.User{}, err
	}

	logUser.Info("User has been created successfully: ", user)
	return user, nil
}

func (u *userRepository) GetUserByUsername(username string) (entity.User, error) {
	var user entity.User

	logUser.Info("Starting to retrive a user by username in the repository layer")

	err := u.db.QueryRow(`SELECT id_user, username, password, role FROM mst_user WHERE username = $1`, username).Scan(&user.Id_user, &user.Username, &user.Password, &user.Role)

	if err != nil {
		logUser.Error("Failed to retrive the user: ", err)
		return entity.User{}, err
	}

	logUser.Info("Getting user by username was successfully: ", user)
	return user, nil
}

func (u *userRepository) GetUserByID(id string) (entity.User, error) {
	var user entity.User

	logUser.Info("Starting to retrive a user by id in the repository layer")

	err := u.db.QueryRow(`SELECT id_user, username, password, role FROM mst_user WHERE id_user = $1`, id).Scan(&user.Id_user, &user.Username, &user.Password, &user.Role)

	if err != nil {
		logUser.Error("Failed to retrive the user: ", err)
		return entity.User{}, err
	}

	logUser.Info("Getting user by id was successfully: ", user)
	return user, nil

}
func (u *userRepository) UpdateUser(user entity.User) (entity.User, error) {
	logUser.Info("Starting to update user in the repository layer")

	_, err := u.db.Exec(`UPDATE mst_user SET username = $1, password = $2, role = $3 WHERE id_user = $4`, user.Username, user.Password, user.Role, user.Id_user)

	if err != nil {
		logUser.Error("Failed to update the user: ", err)
		return entity.User{}, err
	}

	logUser.Info("User has been updated successfully: ", user)
	return user, nil
}
func (u *userRepository) DeleteUser(id string) error {
	logUser.Info("Starting to delete user in the repository layer")

	_, err := u.db.Exec(`DELETE FROM mst_user WHERE id_user = $1`, id)

	if err != nil {
		logUser.Error("Failed to delete the user: ", err)
		return err
	}

	logUser.Info("User has been deleted successfully: ", id)
	return nil
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

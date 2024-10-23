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

type userRepository struct {
	db *sql.DB
}

func (u *userRepository) CreateUser(user entity.User) (entity.User, error) {
	log := logger.GetLogger()
	log.Info("Start creating a new user")

	err := u.db.QueryRow(`INSERT INTO mst_user (username, password, role) VALUES ($1, $2, $3) RETURNING id_user`, user.Username, user.Password, user.Role).Scan(&user.Id_user)

	if err != nil {
		log.Error("Failed to create user: ", err)
		return entity.User{}, err
	}

	log.Info("User created successfully: ", user)
	return user, nil
}

func (u *userRepository) GetUserByUsername(username string) (entity.User, error) {
	var user entity.User

	log := logger.GetLogger()
	log.Info("Start getting user by username")

	err := u.db.QueryRow(`SELECT id_user, username, password, role FROM mst_user WHERE username = $1`, username).Scan(&user.Id_user, &user.Username, &user.Password, &user.Role)

	if err != nil {
		log.Error("Failed to get user: ", err)
		return entity.User{}, err
	}

	log.Info("Get user by username successfully: ", user)
	return user, nil
}

func (u *userRepository) GetUserByID(id string) (entity.User, error) {
	var user entity.User

	log := logger.GetLogger()
	log.Info("Start getting user by id")

	err := u.db.QueryRow(`SELECT id_user, username, password, role FROM mst_user WHERE id_user = $1`, id).Scan(&user.Id_user, &user.Username, &user.Password, &user.Role)

	if err != nil {
		log.Error("Failed to get user: ", err)
		return entity.User{}, err
	}

	log.Info("Get user by id successfully: ", user)
	return user, nil

}
func (u *userRepository) UpdateUser(user entity.User) (entity.User, error) {
	log := logger.GetLogger()
	log.Info("Start updating user")

	_, err := u.db.Exec(`UPDATE mst_user SET username = $1, password = $2, role = $3 WHERE id_user = $4`, user.Username, user.Password, user.Role, user.Id_user)

	if err != nil {
		log.Error("Failed to update user: ", err)
		return entity.User{}, err
	}

	log.Info("User updated successfully: ", user)
	return user, nil
}
func (u *userRepository) DeleteUser(id string) error {
	log := logger.GetLogger()
	log.Info("Start deleting user")

	_, err := u.db.Exec(`DELETE FROM mst_user WHERE id_user = $1`, id)

	if err != nil {
		log.Error("Failed to delete user: ", err)
		return err
	}

	log.Info("User deleted successfully: ", id)
	return nil
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

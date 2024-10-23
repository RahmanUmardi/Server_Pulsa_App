package repository

import (
	"database/sql"
	"server-pulsa-app/internal/entity"
)

type UserRepository interface {
	RegisterUser(user entity.User) (entity.User, error)
	GetUserByID(id string) (entity.User, error)
	GetUserByUsername(username string) (entity.User, error)
	UpdateUser(user entity.User) (entity.User, error)
	DeleteUser(id string) error
}

type userRepository struct {
	db *sql.DB
}

func (u *userRepository) RegisterUser(user entity.User) (entity.User, error) {
	err := u.db.QueryRow(`INSERT INTO mst_user (username, password, role) VALUES ($1, $2, $3) RETURNING id_user`, user.Username, user.Password, user.Role).Scan(&user.Id_user)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (u *userRepository) GetUserByUsername(username string) (entity.User, error) {
	var user entity.User
	err := u.db.QueryRow(`SELECT id_user, username, password, role FROM mst_user WHERE username = $1`, username).Scan(&user.Id_user, &user.Username, &user.Password, &user.Role)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (u *userRepository) GetUserByID(id string) (entity.User, error) {
	var user entity.User

	err := u.db.QueryRow(`SELECT id_user, username, password, role FROM mst_user WHERE id_user = $1`, id).Scan(&user.Id_user, &user.Username, &user.Password, &user.Role)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil

}
func (u *userRepository) UpdateUser(user entity.User) (entity.User, error) {
	_, err := u.db.Exec(`UPDATE mst_user SET username = $1, password = $2, role = $3 WHERE id_user = $4`, user.Username, user.Password, user.Role, user.Id_user)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}
func (u *userRepository) DeleteUser(id string) error {
	_, err := u.db.Exec(`DELETE FROM mst_user WHERE id_user = $1`, id)
	if err != nil {
		return err
	}
	return nil
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

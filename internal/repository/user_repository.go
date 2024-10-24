package repository

import (
	"database/sql"
	"log"
	"server-pulsa-app/internal/entity"
	"strings"
)

type UserRepository interface {
	CreateUser(user entity.User) (entity.User, error)
	ListUser() ([]entity.User, error)
	GetUserByID(id string) (entity.User, error)
	GetUserByUsername(username string) (entity.User, error)
	UpdateUser(user, payload entity.User) (entity.User, error)
	DeleteUser(id string) error
}

type userRepository struct {
	db *sql.DB
}

func (u *userRepository) CreateUser(user entity.User) (entity.User, error) {
	err := u.db.QueryRow(`INSERT INTO mst_user (username, password, role) VALUES ($1, $2, $3) RETURNING id_user`, user.Username, user.Password, user.Role).Scan(&user.Id_user)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (u *userRepository) ListUser() ([]entity.User, error) {
	var users []entity.User

	rows, err := u.db.Query(`SELECT id_user, username, password, role FROM mst_user`)
	if err != nil {
		log.Printf("UserRepository.ListUser: %v \n", err.Error())
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.Id_user, &user.Username, &user.Password, &user.Role)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
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
func (u *userRepository) UpdateUser(user, payload entity.User) (entity.User, error) {
	if strings.TrimSpace(payload.Username) != "" {
		user.Username = payload.Username
	}
	if strings.TrimSpace(payload.Password) != "" {
		user.Password = payload.Password
	}
	if strings.TrimSpace(payload.Role) != "" {
		user.Role = payload.Role
	}
	_, err := u.db.Exec(`UPDATE mst_user SET username = $2, password = $3, role = $4 WHERE id_user = $1`, user.Id_user, user.Username, user.Password, user.Role)
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

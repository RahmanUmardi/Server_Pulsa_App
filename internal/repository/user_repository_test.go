package repository

import (
	"testing"

	"server-pulsa-app/internal/entity"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' occurred when opening a mock database connection", err)
	}
	defer db.Close()

	repo := NewUserRepository(db)

	user := entity.User{
		Id_user:  "1",
		Username: "test",
		Password: "test",
		Role:     "test",
	}

	mock.ExpectQuery("INSERT INTO mst_user").
		WithArgs(user.Id_user, user.Username, user.Password, user.Role).
		WillReturnRows(sqlmock.NewRows([]string{"id_user"}).AddRow("1"))

	createdUser, err := repo.CreateUser(user)
	if err != nil {
		t.Errorf("error was not expected while creating user: %s", err)
	}

	if createdUser.Id_user != "1" {
		t.Errorf("expected  Id_user to be '1', got '%s'", createdUser.Id_user)
	}
}

func TestGetByUsername(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' occurred when opening a mock database connection", err)
	}
	defer db.Close()

	repo := NewUserRepository(db)

	Username := "test"
	expectedUser := entity.User{
		Id_user:  "1",
		Username: "test",
		Password: "test",
		Role:     "test",
	}

	mock.ExpectQuery("SELECT id_user, username, password, role FROM mst_user WHERE username= ?").
		WithArgs(Username).
		WillReturnRows(sqlmock.NewRows([]string{"id_user", "username", "password", "role"}).AddRow(expectedUser.Id_user, expectedUser.Username, expectedUser.Password, expectedUser.Role))

	user, err := repo.GetUserByUsername(Username)
	if err != nil {
		t.Errorf("error was not expected while getting user: %s", err)
	}

	if user != expectedUser {
		t.Errorf("expected user %v, got %v", expectedUser, user)
	}
}

func TestGetById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' occurred when opening a mock database connection", err)
	}
	defer db.Close()

	repo := NewUserRepository(db)

	Id_user := "1"
	expectedUser := entity.User{
		Id_user:  "1",
		Username: "test",
		Password: "test",
		Role:     "test",
	}

	mock.ExpectQuery("SELECT id_user, username, password, role FROM mst_user WHERE id_user= ?").
		WithArgs(Id_user).
		WillReturnRows(sqlmock.NewRows([]string{"id_user", "username", "password", "role"}).
			AddRow(expectedUser.Id_user, expectedUser.Username, expectedUser.Password, expectedUser.Role))

	user, err := repo.GetUserByID(Id_user)
	if err != nil {
		t.Errorf("error was not expected while getting user: %s", err)
	}

	if user != expectedUser {
		t.Errorf("expected user %v, got %v", expectedUser, user)
	}
}

func TestList(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' occurred when opening a mock database connection", err)
	}
	defer db.Close()

	repo := NewUserRepository(db)

	mock.ExpectQuery("SELECT * FROM mst_user").
		WillReturnRows(sqlmock.NewRows([]string{"id_user", "username", "password", "role"}).
			AddRow("1", "userA", "passA", "admin").
			AddRow("2", "userB", "passB", "user"))

	users, err := repo.ListUser()
	if err != nil {
		t.Errorf("error was not expected while listing users: %s", err)
	}

	if len(users) != 2 {
		t.Errorf("expected 2 users, got %d", len(users))
	}
}

func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' occurred when opening a mock database connection", err)
	}
	defer db.Close()

	repo := NewUserRepository(db)

	Id_user := "1"
	user := entity.User{
		Username: "test",
		Password: "test",
		Role:     "test",
	}

	mock.ExpectExec("UPDATE mst_product SET username = ?, password = ?, role = ? WHERE id_user = ?").
		WithArgs(user.Username, user.Password, user.Role, Id_user).
		WillReturnResult(sqlmock.NewResult(1, 1))

	updatedUser, err := repo.UpdateUser(user, user)
	if err != nil {
		t.Errorf("error was not expected while updating user: %s", err)
	}

	if updatedUser.Id_user != Id_user {
		t.Errorf("expected Id_user to be '%s', got '%s'", Id_user, updatedUser.Id_user)
	}
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' occurred when opening a mock database connection", err)
	}
	defer db.Close()

	repo := NewUserRepository(db)

	Id_user := "1"

	mock.ExpectExec("DELETE FROM mst_user WHERE id_user = ?").
		WithArgs(Id_user).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.DeleteUser(Id_user)
	if err != nil {
		t.Errorf("error was not expected while deleting user: %s", err)
	}
}

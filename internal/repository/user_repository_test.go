package repository

import (
	"database/sql"
	"regexp"
	"server-pulsa-app/internal/entity"
	"server-pulsa-app/internal/logger"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
)

var expectedUser = entity.User{
	Id_user:  "uuid-user-test",
	Username: "username-test",
	Password: "password-test",
	Role:     "test",
}

type userRepositoryTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	mockSql sqlmock.Sqlmock
	ur      UserRepository
	log     logger.Logger
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(userRepositoryTestSuite))
}

func (u *userRepositoryTestSuite) SetupTest() {
	mockDb, mockSql, err := sqlmock.New()
	u.NoError(err)

	u.mockDb = mockDb
	u.mockSql = mockSql
	u.log = logger.NewLogger()
	u.ur = NewUserRepository(mockDb, &u.log)
}

func (u *userRepositoryTestSuite) TestCreate_success() {
	u.mockSql.ExpectQuery(regexp.QuoteMeta("INSERT INTO mst_user (username, password, role) VALUES ($1, $2, $3) RETURNING id_user")).WillReturnRows(
		sqlmock.NewRows([]string{"id_user"}).AddRow(expectedUser.Id_user),
	)

	_, err := u.ur.CreateUser(expectedUser)

	u.Nil(err)
}

func (u *userRepositoryTestSuite) TestCreate_fail() {
	u.mockSql.ExpectQuery(regexp.QuoteMeta("INSERT INTO mst_user (id_user, username, password, role) VALUES ($1, $2, $3, $4) RETURNING id_user")).WillReturnError(sql.ErrNoRows)

	_, err := u.ur.CreateUser(expectedUser)

	u.NotNil(err)
}
func (u *userRepositoryTestSuite) TestGetId_success() {

	userRows := sqlmock.NewRows([]string{"id_user", "username", "password", "role"}).AddRow(
		expectedUser.Id_user,
		expectedUser.Username,
		expectedUser.Password,
		expectedUser.Role,
	)

	u.mockSql.ExpectQuery(regexp.QuoteMeta("SELECT id_user, username, password, role FROM mst_user WHERE id_user = $1")).
		WithArgs(expectedUser.Id_user).WillReturnRows(
		userRows,
	)

	user, err := u.ur.GetUserByID("uuid-user-test")

	u.Nil(err)
	u.Equal(expectedUser, user)
}

func (u *userRepositoryTestSuite) TestGetId_fail() {
	u.mockSql.ExpectQuery(regexp.QuoteMeta("SELECT id_user, username, password, role FROM mst_user WHERE id_user = $1")).
		WithArgs(expectedUser.Id_user).WillReturnError(sql.ErrNoRows)

	_, err := u.ur.GetUserByID("uuid-merchant-test")

	u.NotNil(err)
}

func (u *userRepositoryTestSuite) TestGetUsername_success() {

	userRows := sqlmock.NewRows([]string{"id_user", "username", "password", "role"}).AddRow(
		expectedUser.Id_user,
		expectedUser.Username,
		expectedUser.Password,
		expectedUser.Role,
	)

	u.mockSql.ExpectQuery(regexp.QuoteMeta("SELECT id_user, username, password, role FROM mst_user WHERE username = $1")).
		WithArgs(expectedUser.Username).WillReturnRows(
		userRows,
	)

	user, err := u.ur.GetUserByUsername("username-test")

	u.Nil(err)
	u.Equal(expectedUser, user)
}

func (u *userRepositoryTestSuite) TestGetUsername_fail() {
	u.mockSql.ExpectQuery(regexp.QuoteMeta("SELECT id_user, username, password, role FROM mst_user WHERE username = $2")).
		WithArgs(expectedUser.Username).WillReturnError(sql.ErrNoRows)

	_, err := u.ur.GetUserByUsername("username-test")

	u.NotNil(err)
}

func (u *userRepositoryTestSuite) TestList_success() {
	userRows := sqlmock.NewRows([]string{"id_user", "username", "password", "role"}).AddRow(
		expectedUser.Id_user,
		expectedUser.Username,
		expectedUser.Password,
		expectedUser.Role,
	)

	u.mockSql.ExpectQuery(regexp.QuoteMeta("SELECT id_user, username, password, role FROM mst_user")).WillReturnRows(
		userRows,
	)

	users, err := u.ur.ListUser()

	u.Nil(err)
	u.Equal([]entity.User{expectedUser}, users)
}

func (u *userRepositoryTestSuite) TestList_fail() {
	u.mockSql.ExpectQuery(regexp.QuoteMeta("SELECT id_user, username, password, role FROM mst_user")).WillReturnError(sql.ErrNoRows)

	_, err := u.ur.ListUser()

	u.NotNil(err)
}

func (u *userRepositoryTestSuite) TestDelete_fail() {
	u.mockSql.ExpectQuery(regexp.QuoteMeta("DELETE FROM mst_user WHERE id_user = $1")).WillReturnError(sql.ErrNoRows)

	err := u.ur.DeleteUser(expectedUser.Id_user)

	u.NotNil(err)
}

func (u *userRepositoryTestSuite) TestUpdate_fail() {
	user := entity.User{
		Id_user:  "uuid-user-test",
		Username: "username-test",
		Password: "password-test",
		Role:     "test",
	}

	u.mockSql.ExpectQuery(regexp.QuoteMeta("UPDATE mst_merchant SET username = $1, password = $2, role = $3 WHERE id_user = $4")).WillReturnError(sql.ErrNoRows)

	_, err := u.ur.UpdateUser(user, expectedUser)

	u.NotNil(err)
}

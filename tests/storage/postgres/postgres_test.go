package postgres

import (
	"authentication-service/models"
	"authentication-service/storage/postgres"
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserRepository_GetUserByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	userRepo := postgres.NewUserRepository(sqlxDB)

	rows := sqlmock.NewRows([]string{"id", "full_name", "username", "email", "bio", "user_type", "password_hash"}).
		AddRow("1", "Test User", "testuser", "test@example.com", "Bio", "user", "hashedpassword")

	mock.ExpectQuery("SELECT \\* FROM users WHERE id = \\$1").
		WithArgs("1").
		WillReturnRows(rows)

	user, err := userRepo.GetUserByID(context.Background(), "1")
	require.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "testuser", user.Username)
}

func TestUserRepository_UpdateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	userRepo := postgres.NewUserRepository(sqlxDB)

	mock.ExpectExec("UPDATE users SET full_name=\\$1, username=\\$2, email=\\$3, bio=\\$4, user_type=\\$5 WHERE id=\\$6").
		WithArgs("Updated User", "updateduser", "updated@example.com", "Updated Bio", "user", "1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	user := &models.User{
		ID:           "1",
		FullName:     "Updated User",
		Username:     "updateduser",
		Email:        "updated@example.com",
		Bio:          "Updated Bio",
		UserType:     "user",
		PasswordHash: "hashedpassword",
	}

	err = userRepo.UpdateUser(context.Background(), user)
	require.NoError(t, err)
}

func TestUserRepository_DeleteUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	userRepo := postgres.NewUserRepository(sqlxDB)

	mock.ExpectExec("DELETE FROM users WHERE id=\\$1").
		WithArgs("1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = userRepo.DeleteUser(context.Background(), "1")
	require.NoError(t, err)
}

func TestUserRepository_GetUsersInfo(t *testing.T) {

}

func TestUserRepository_GetUserByUsernameOrEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	userRepo := postgres.NewUserRepository(sqlxDB)

	rows := sqlmock.NewRows([]string{"id", "full_name", "username", "email", "bio", "user_type", "password_hash"}).
		AddRow("1", "Test User", "testuser", "test@example.com", "Bio", "user", "hashedpassword")

	mock.ExpectQuery("SELECT \\* FROM users WHERE username = \\$1").
		WithArgs("testuser").
		WillReturnRows(rows)

	user, err := userRepo.GetUserByUsernameOrEmail(context.Background(), "testuser", "")
	require.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "testuser", user.Username)

	rows = sqlmock.NewRows([]string{"id", "full_name", "username", "email", "bio", "user_type", "password_hash"}).
		AddRow("2", "Another User", "anotheruser", "another@example.com", "Another Bio", "user", "hashedpassword")

	mock.ExpectQuery("SELECT \\* FROM users WHERE email = \\$1").
		WithArgs("another@example.com").
		WillReturnRows(rows)

	user, err = userRepo.GetUserByUsernameOrEmail(context.Background(), "", "another@example.com")
	require.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "anotheruser", user.Username)
}

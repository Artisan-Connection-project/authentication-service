package postgres_test

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

func TestAuthRepo_Register(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	userRepo := postgres.NewUserRepository(sqlxDB)
	hasher := postgres.NewBcryptHasher(10)
	authRepo := postgres.NewAuthenticationRepository(userRepo, hasher, sqlxDB)

	user := &models.User{
		FullName:     "Test User",
		Username:     "testuser",
		Email:        "test@example.com",
		Bio:          "Bio",
		UserType:     "user",
		PasswordHash: "plainpassword",
	}

	mock.ExpectExec(`INSERT INTO users \(full_name, username, email, bio, user_type, password_hash\) VALUES \(\$1, \$2, \$3, \$4, \$5, \$6\)`).
		WithArgs(user.FullName, user.Username, user.Email, user.Bio, user.UserType, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = authRepo.Register(context.Background(), user)
	require.NoError(t, err)
	assert.NotEmpty(t, user.PasswordHash)
}

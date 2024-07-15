package postgres

import (
	auth "authentication-service/genproto/authentication_service"
	"authentication-service/models"
	"context"
	"database/sql"
	"fmt"

	"github.com/go-openapi/errors"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id string) (*models.User, error)
	UpdateUser(ctx context.Context, updatedUser *models.User) (*auth.UpdateUserInfoResponse, error)
	DeleteUser(context.Context, string) error
	GetUsersInfo(context.Context, int, int, string) ([]*auth.User, error)
	GetUserByUsernameOrEmail(ctx context.Context, email string, username string) (*models.User, error)
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	err := r.db.GetContext(ctx, &user, "SELECT id, username, email, full_name, user_type, bio, created_at, updated_at FROM users WHERE id = $1 AND deleted_at IS NULL", id)
	if err == sql.ErrNoRows {
		return nil, errors.NotFound("user does not exist")
	}
	return &user, err
}

func (r *userRepository) UpdateUser(ctx context.Context, user *models.User) (*auth.UpdateUserInfoResponse, error) {
	if user.Username == "" {
		return nil, fmt.Errorf("username cannot be empty")
	}
	if user.Email == "" {
		return nil, fmt.Errorf("email cannot be empty")
	}
	if err := uuid.Validate(user.ID); err != nil {
		return nil, fmt.Errorf("invalid user ID")
	}

	query := `
		UPDATE users SET full_name=$1, username=$2, email=$3, bio=$4, user_type=$5, updated_at = NOW() 
		WHERE id=$6 RETURNING 
		id,
		username,
        email,
        bio,
        user_type,
		created_at,
        updated_at
	`
	row := r.db.QueryRowContext(ctx, query, user.FullName, user.Username, user.Email, user.Bio, user.UserType, user.ID)
	if row.Err() == sql.ErrNoRows {
		return nil, errors.NotFound("user does not exist")
	}
	var updatedUser auth.UpdateUserInfoResponse
	if err := row.Scan(&updatedUser.Id,
		&updatedUser.Username,
		&updatedUser.Email,
		&updatedUser.Bio,
		&updatedUser.UserType,
		&updatedUser.CreatedAt,
		&updatedUser.UpdatedAt); err != nil {
		return nil, err
	}
	return &updatedUser, nil
}

func (r *userRepository) DeleteUser(ctx context.Context, id string) error {
	res, err := r.db.ExecContext(ctx, "UPDATE users SET deleted_at=NOW() WHERE id=$1 AND DELETED_AT IS NULL", id)
	if n, _ := res.RowsAffected(); n == 0 {
		return fmt.Errorf("user does not exist")
	}
	return err
}

func (r *userRepository) GetUsersInfo(ctx context.Context, page, limit int, orderBy string) ([]*auth.User, error) {
	if orderBy == "" {
		orderBy = "username"
	}

	query := `
		SELECT id, username, full_name, bio, user_type, email
		FROM users
		WHERE deleted_at IS NULL
		ORDER BY ` + orderBy + `
		LIMIT $1 OFFSET $2
	`

	offset := (page - 1) * limit
	var users []*auth.User

	err := r.db.SelectContext(ctx, &users, query, limit, offset)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) GetUserByUsernameOrEmail(ctx context.Context, email string, username string) (*models.User, error) {
	var user models.User
	if username != "" {
		err := r.db.GetContext(ctx, &user, "SELECT id, username, full_name, bio, user_type, email, password_hash FROM users WHERE username = $1 AND DELETED_AT IS NULL", username)
		if err != nil {
			return nil, err
		}
		return &user, nil
	}
	if email != "" {
		err := r.db.GetContext(ctx, &user, "SELECT id, username, full_name, bio, user_type, email, password_hash FROM users WHERE email = $1 AND DELETED_AT IS NULL", email)
		if err != nil {
			return nil, err
		}
		return &user, nil
	}
	return nil, fmt.Errorf("user %s not found", username)
}

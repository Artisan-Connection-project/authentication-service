package postgres

import (
	"authentication-service/models"
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id string) (*models.User, error)
	UpdateUser(ctx context.Context, updatedUser *models.User) error
	DeleteUser(context.Context, string) error
	GetUsersInfo(context.Context, int, int, string) ([]models.User, error)
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
	err := r.db.GetContext(ctx, &user, "SELECT id, username, email, full_name, user_type, bio, created_at FROM users WHERE id = $1 AND deleted_at IS NULL", id)
	return &user, err
}

func (r *userRepository) UpdateUser(ctx context.Context, user *models.User) error {
	if user.Username == "" {
		return fmt.Errorf("username cannot be empty")
	}
	if user.Email == "" {
		return fmt.Errorf("email cannot be empty")
	}
	if err := uuid.Validate(user.ID); err != nil {
		return fmt.Errorf("invalid user ID")
	}

	row, err := r.db.ExecContext(ctx, "UPDATE users SET full_name=$1, username=$2, email=$3, bio=$4, user_type=$5, updated_at = NOW() WHERE id=$6 ",
		user.FullName, user.Username, user.Email, user.Bio, user.UserType, user.ID)
	if n, _ := row.RowsAffected(); n != 0 {
		return fmt.Errorf("no users found")
	}
	return err
}

func (r *userRepository) DeleteUser(ctx context.Context, id string) error {
	res, err := r.db.ExecContext(ctx, "UPDATE users SET deleted_at=NOW() WHERE id=$1 AND DELETED_AT IS NULL", id)
	if n, _ := res.RowsAffected(); n == 0 {
		return fmt.Errorf("no such user")
	}
	return err
}

func (r *userRepository) GetUsersInfo(ctx context.Context, page, limit int, orderBy string) ([]models.User, error) {
	offset := (page - 1) * limit
	var users []models.User
	err := r.db.SelectContext(ctx, &users, "SELECT id, username, full_name, bio, user_type, email FROM users ORDER BY $3 OFFSET $1 LIMIT $2 AND deleted_at IS NULL", offset, limit, orderBy)
	return users, err
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

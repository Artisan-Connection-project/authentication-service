package postgres

import (
	"authentication-service/models"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type AuthenticationRepository interface {
	Login(ctx context.Context, username, email, password string) (*models.User, error)
	Logout(ctx context.Context, token string) error
	Register(ctx context.Context, user *models.User) error
	ResetPassword(ctx context.Context, email, username, newPassword string) error
	ChangePassword(ctx context.Context, userID, currentPassword, newPassword string) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	UpdatePassword(ctx context.Context, email, newPassword string) error
	UpdateUserToActive(ctx context.Context, email string) error
}

type authenticationRepositoryImpl struct {
	userRepository UserRepository
	hasher         Hasher
	db             *sqlx.DB
}

func NewAuthenticationRepository(userRepository UserRepository, hasher Hasher, db *sqlx.DB) AuthenticationRepository {
	return &authenticationRepositoryImpl{userRepository: userRepository, hasher: hasher, db: db}
}

// Implement the GetUserByEmail method
func (r *authenticationRepositoryImpl) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.db.GetContext(ctx, &user, "SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Implement the UpdatePassword method
func (r *authenticationRepositoryImpl) UpdatePassword(ctx context.Context, email, newPassword string) error {
	hashedPassword := r.hasher.Hash(newPassword)
	_, err := r.db.ExecContext(ctx, "UPDATE users SET password_hash = $1 WHERE email = $2", hashedPassword, email)
	return err
}

func (r *authenticationRepositoryImpl) Register(ctx context.Context, user *models.User) error {
	hashedPassword := r.hasher.Hash(user.PasswordHash)

	_, err := r.db.ExecContext(ctx, `
        INSERT INTO users (full_name, username, email, bio, user_type, password_hash, deleted_at)
        VALUES ($1, $2, $3, $4, $5, $6, now())
    `, user.FullName, user.Username, user.Email, user.Bio, user.UserType, hashedPassword)
	return err
}

func (r *authenticationRepositoryImpl) Login(ctx context.Context, username, email, password string) (*models.User, error) {
	user, err := r.userRepository.GetUserByUsernameOrEmail(ctx, email, username)
	if err != nil {
		return nil, err
	}
	if !r.hasher.Compare(user.PasswordHash, password) {
		return nil, fmt.Errorf("password mismatch")
	}
	return user, nil
}

func (r *authenticationRepositoryImpl) Logout(ctx context.Context, token string) error {
	// Implement logout logic using the provided token
	return nil
}
func (r *authenticationRepositoryImpl) ResetPassword(ctx context.Context, email string, username string, newPassword string) error {
	var user models.User
	if username != "" {
		err := r.db.GetContext(ctx, &user, "SELECT * FROM users WHERE username = $1", username)
		if err != nil {
			return err
		}
	}
	if email != "" {
		err := r.db.GetContext(ctx, &user, "SELECT * FROM users WHERE email = $1", email)
		if err != nil {
			return err
		}
	}

	hashedPassword := r.hasher.Hash(newPassword)

	user.PasswordHash = hashedPassword

	_, err := r.userRepository.UpdateUser(ctx, &user)
	if err != nil {
		return err
	}

	return nil
}

func (r *authenticationRepositoryImpl) ChangePassword(ctx context.Context, email, currentPassword, newPassword string) error {
	user, err := r.userRepository.GetUserByUsernameOrEmail(ctx, email, "")
	if err != nil {
		return err
	}

	if !r.hasher.Compare(user.PasswordHash, currentPassword) {
		return fmt.Errorf("Current password mismatch")
	}

	hashedPassword := r.hasher.Hash(newPassword)
	user.PasswordHash = hashedPassword

	_, err = r.db.ExecContext(ctx, `
        UPDATE users
        SET password_hash = $1
        WHERE email = $2
    `, user.PasswordHash, user.Email)
	return err
}

func (r *authenticationRepositoryImpl) UpdateUserToActive(ctx context.Context, email string) error {
	query := `
		UPDATE users SET deleted_at = NULL WHERE email = $1
	`
	_, err := r.db.ExecContext(ctx, query, email)
	if err != nil {
		return err
	}
	return nil
}

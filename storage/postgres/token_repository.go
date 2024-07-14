package postgres

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type TokenRepository interface {
	CreateToken(ctx context.Context, userID string) (string, error)
	DeleteToken(ctx context.Context, token string) error
	IsTokenValid(ctx context.Context, token string) bool
	GetUserIDByToken(ctx context.Context, token string) (string, error)
	GetRefreshTokensByUserID(ctx context.Context, userID string) ([]string, error)
	DeleteRefreshTokensByUserID(ctx context.Context, userID string) error
	UpdateRefreshToken(ctx context.Context, userID, token string) error
	DeleteExpiredRefreshTokens(ctx context.Context) error
	CreateRefreshToken(ctx context.Context, userID, token string) error
	IsRefreshTokenValid(ctx context.Context, userID, token string) bool
	GetRefreshToken(ctx context.Context, userID, token string) (string, error)
	DeleteRefreshTokenByUserIDAndToken(ctx context.Context, userID, token string) error
}

type tokenRepositoryImpl struct {
	db *sqlx.DB
}

func NewTokenRepository(db *sqlx.DB) TokenRepository {
	return &tokenRepositoryImpl{db: db}
}

func (r *tokenRepositoryImpl) CreateToken(ctx context.Context, userID string) (string, error) {
	// Implement token creation logic here
	return "", nil
}

func (r *tokenRepositoryImpl) DeleteToken(ctx context.Context, token string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM refresh_tokens WHERE token = $1", token)
	return err
}

func (r *tokenRepositoryImpl) IsTokenValid(ctx context.Context, token string) bool {
	// Implement token validation logic here
	return false
}

func (r *tokenRepositoryImpl) GetUserIDByToken(ctx context.Context, token string) (string, error) {
	var userID string
	err := r.db.GetContext(ctx, &userID, "SELECT user_id FROM refresh_tokens WHERE token = $1", token)
	return userID, err
}

func (r *tokenRepositoryImpl) GetRefreshTokensByUserID(ctx context.Context, userID string) ([]string, error) {
	var tokens []string
	err := r.db.SelectContext(ctx, &tokens, "SELECT token FROM refresh_tokens WHERE user_id = $1", userID)
	return tokens, err
}

func (r *tokenRepositoryImpl) DeleteRefreshTokensByUserID(ctx context.Context, userID string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM refresh_tokens WHERE user_id = $1", userID)
	return err
}

func (r *tokenRepositoryImpl) UpdateRefreshToken(ctx context.Context, userID, token string) error {
	_, err := r.db.ExecContext(ctx, "UPDATE refresh_tokens SET last_used_at = NOW() WHERE user_id = $1 AND token = $2", userID, token)
	return err
}

func (r *tokenRepositoryImpl) DeleteExpiredRefreshTokens(ctx context.Context) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM refresh_tokens WHERE last_used_at < NOW() - INTERVAL '1 day'")
	return err
}

func (r *tokenRepositoryImpl) CreateRefreshToken(ctx context.Context, userID, token string) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO refresh_tokens (user_id, token, last_used_at) VALUES ($1, $2, NOW())", userID, token)
	return err
}

func (r *tokenRepositoryImpl) IsRefreshTokenValid(ctx context.Context, userID, token string) bool {
	var count int
	err := r.db.GetContext(ctx, &count, "SELECT COUNT(*) FROM refresh_tokens WHERE user_id = $1 AND token = $2 AND last_used_at > NOW() - INTERVAL '1 day'", userID, token)
	return err == nil && count > 0
}

func (r *tokenRepositoryImpl) GetRefreshToken(ctx context.Context, userID, token string) (string, error) {
	var refreshToken string
	err := r.db.GetContext(ctx, &refreshToken, "SELECT token FROM refresh_tokens WHERE user_id = $1 AND token = $2", userID, token)
	return refreshToken, err
}

func (r *tokenRepositoryImpl) DeleteRefreshTokenByUserIDAndToken(ctx context.Context, userID, token string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM refresh_tokens WHERE user_id = $1 AND token = $2", userID, token)
	return err
}

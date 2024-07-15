package postgres

import (
	"context"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type TokenRepository interface {
	CreateRefreshToken(ctx context.Context, email string, token string) error
	DeleteToken(ctx context.Context, token string) error
	GetTokenByEmail(ctx context.Context, email string) (string, error) // Example method to get token by email
}

type tokenRepositoryImpl struct {
	db *sqlx.DB
}

func NewTokenRepository(db *sqlx.DB) TokenRepository {
	return &tokenRepositoryImpl{db: db}
}

func (r *tokenRepositoryImpl) CreateRefreshToken(ctx context.Context, email string, token string) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO refresh_tokens (email, token) VALUES ($1, $2)", email, token)
	return err
}

func (r *tokenRepositoryImpl) DeleteToken(ctx context.Context, token string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM refresh_tokens WHERE token = $1", token)
	return err
}

func (r *tokenRepositoryImpl) GetTokenByEmail(ctx context.Context, email string) (string, error) {
	var token string
	err := r.db.QueryRowContext(ctx, "SELECT token FROM refresh_tokens WHERE email = $1", email).Scan(&token)
	if err != nil {
		return "", err
	}
	return token, nil
}

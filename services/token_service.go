package services

import (
	auth "authentication-service/genproto/authentication_service"
	"authentication-service/storage/postgres"
	"context"
)

type TokenService interface {
	RefreshToken(ctx context.Context, req *auth.RefreshTokenRequest) (*auth.RefreshTokenResponse, error)
	VerifyToken(ctx context.Context, req *auth.VerifyTokenRequest) (*auth.VerifyTokenResponse, error)
	CancelToken(ctx context.Context, req *auth.CancelTokenRequest) (*auth.CancelTokenResponse, error)
}

type tokenServiceImpl struct {
	auth.UnimplementedAuthenticationServiceServer
	tokenRepo postgres.TokenRepository
}

func NewTokenService(tokenRepo postgres.TokenRepository) TokenService {
	return &tokenServiceImpl{
		tokenRepo: tokenRepo,
	}
}

func (s *tokenServiceImpl) RefreshToken(ctx context.Context, req *auth.RefreshTokenRequest) (*auth.RefreshTokenResponse, error) {
	// Implement token refresh logic
	return &auth.RefreshTokenResponse{}, nil
}

func (s *tokenServiceImpl) VerifyToken(ctx context.Context, req *auth.VerifyTokenRequest) (*auth.VerifyTokenResponse, error) {
	// Implement token verification logic
	return &auth.VerifyTokenResponse{}, nil
}

func (s *tokenServiceImpl) CancelToken(ctx context.Context, req *auth.CancelTokenRequest) (*auth.CancelTokenResponse, error) {
	// Implement token cancellation logic
	return &auth.CancelTokenResponse{}, nil
}

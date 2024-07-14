package services

import (
	auth "authentication-service/genproto/authentication_service"
	"authentication-service/storage/postgres"
	"context"
	"time"

	"github.com/golang-jwt/jwt"
)

type claims struct {
	jwt.StandardClaims
	Email    string `json:"email,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type TokenService interface {
	RefreshToken(ctx context.Context, req *auth.RefreshTokenRequest) (*auth.RefreshTokenResponse, error)
	VerifyToken(ctx context.Context, req *auth.VerifyTokenRequest) (*auth.VerifyTokenResponse, error)
	CancelToken(ctx context.Context, req *auth.CancelTokenRequest) (*auth.CancelTokenResponse, error)
	GenerateToken(ctx context.Context, req *auth.GenerateTokenRequest) (*auth.GenerateTokenResponse, error)
}

type tokenServiceImpl struct {
	jwtSecretKey []byte
	auth.UnimplementedAuthenticationServiceServer
	tokenRepo postgres.TokenRepository
}

func NewTokenService(tokenRepo postgres.TokenRepository, jwtSecretKey string) TokenService {
	return &tokenServiceImpl{
		jwtSecretKey: []byte(jwtSecretKey),
		tokenRepo:    tokenRepo,
	}
}

func (s *tokenServiceImpl) VerifyToken(ctx context.Context, req *auth.VerifyTokenRequest) (*auth.VerifyTokenResponse, error) {
	// Implement token verification logic
	return &auth.VerifyTokenResponse{}, nil
}

func (s *tokenServiceImpl) CancelToken(ctx context.Context, req *auth.CancelTokenRequest) (*auth.CancelTokenResponse, error) {
	// Implement token cancellation logic
	return &auth.CancelTokenResponse{}, nil
}

func (s *tokenServiceImpl) GenerateToken(ctx context.Context, req *auth.GenerateTokenRequest) (*auth.GenerateTokenResponse, error) {
	expirationTimeForRefreshToken := time.Now().Add(time.Hour * 48)
	expirationTimeForAccessToken := time.Now().Add(time.Minute * 30)

	claimsForRefreshToken := &claims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "authentication-service",
			ExpiresAt: expirationTimeForRefreshToken.Unix(),
		},
		Email:    req.Email,
		Username: req.Username,
		Password: req.Password,
	}

	claimsForAccessToken := &claims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "authentication-service",
			ExpiresAt: expirationTimeForAccessToken.Unix(),
		},
		Email:    req.Email,
		Username: req.Username,
		Password: req.Password,
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsForAccessToken)
	accessTokenString, err := accessToken.SignedString(s.jwtSecretKey)
	if err != nil {
		return nil, err
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsForRefreshToken)
	refreshTokenString, err := refreshToken.SignedString(s.jwtSecretKey)
	if err != nil {
		return nil, err
	}

	err = s.tokenRepo.InsertRefreshToken(ctx, req.Email, refreshTokenString)
	if err != nil {
		return nil, err
	}
	return &auth.GenerateTokenResponse{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}

func (s *tokenServiceImpl) RefreshToken(ctx context.Context, req *auth.RefreshTokenRequest) (*auth.RefreshTokenResponse, error) {
	token, err := jwt.ParseWithClaims(req.RefreshToken, &claims{}, func(token *jwt.Token) (interface{}, error) {
		return s.jwtSecretKey, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(*claims)
	if !ok || claims.ExpiresAt < time.Now().Unix() {
		return nil, err
	}

	expirationTimeForAccessToken := time.Now().Add(time.Minute * 30)
	claims.StandardClaims.ExpiresAt = expirationTimeForAccessToken.Unix()

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessTokenString, err := accessToken.SignedString(s.jwtSecretKey)
	if err != nil {
		return nil, err
	}

	return &auth.RefreshTokenResponse{
		AccessToken: accessTokenString,
	}, nil
}

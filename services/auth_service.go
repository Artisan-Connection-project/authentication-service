package services

import (
	auth "authentication-service/genproto/authentication_service"
	"authentication-service/storage/postgres"
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type AuthenticationService interface {
	Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error)
	Logout(ctx context.Context, req *auth.LogoutRequest) (*auth.LogoutResponse, error)
	Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error)
	ResetPassword(ctx context.Context, req *auth.ResetPasswordRequest) (*auth.ResetPasswordResponse, error)
	ChangePassword(ctx context.Context, req *auth.ChangePasswordRequest) (*auth.ChangePasswordResponse, error)
}

type authenticationServiceImpl struct {
	auth.UnimplementedAuthenticationServiceServer
	authRepo     postgres.AuthenticationRepository
	tokenService TokenService
	redisClient  *redis.Client
	emailSender  EmailSender
}

func NewAuthenticationService(userRepo postgres.UserRepository, authRepo postgres.AuthenticationRepository, tokenService TokenService, redisClient *redis.Client, emailSender EmailSender) AuthenticationService {
	return &authenticationServiceImpl{
		authRepo:     authRepo,
		tokenService: tokenService,
		redisClient:  redisClient,
		emailSender:  emailSender,
	}
}

func (s *authenticationServiceImpl) ResetPassword(ctx context.Context, req *auth.ResetPasswordRequest) (*auth.ResetPasswordResponse, error) {
	user, err := s.authRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	resetToken := generateRandomToken() // Implement this function to generate a random token
	err = s.redisClient.Set(ctx, resetToken, user.Email, time.Hour).Err()
	if err != nil {
		return nil, err
	}

	err = s.emailSender.SendResetEmail(user.Email, resetToken) // Implement this function to send email
	if err != nil {
		return nil, err
	}

	return &auth.ResetPasswordResponse{}, nil
}

func (s *authenticationServiceImpl) ChangePassword(ctx context.Context, req *auth.ChangePasswordRequest) (*auth.ChangePasswordResponse, error) {
	user, err := s.authRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if !s.authRepo.Hasher.Compare(user.PasswordHash, req.CurrentPassword) {
		return nil, fmt.Errorf("current password mismatch")
	}

	hashedPassword := s.authRepo.Hasher.Hash(req.NewPassword)
	user.PasswordHash = hashedPassword

	err = s.authRepo.UpdatePassword(ctx, user)
	if err != nil {
		return nil, err
	}

	return &auth.ChangePasswordResponse{}, nil
}

func (s *authenticationServiceImpl) Logout(ctx context.Context, req *auth.LogoutRequest) (*auth.LogoutResponse, error) {
	err := s.tokenService.InvalidateToken(ctx, req.Token) // Implement this function to invalidate token
	if err != nil {
		return nil, err
	}
	return &auth.LogoutResponse{}, nil
}

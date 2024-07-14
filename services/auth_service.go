package services

import (
	auth "authentication-service/genproto/authentication_service"
	"authentication-service/storage/postgres"
	"context"
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
	authRepo postgres.AuthenticationRepository
}

func NewAuthenticationService(userRepo postgres.UserRepository, authRepo postgres.AuthenticationRepository) AuthenticationService {
	return &authenticationServiceImpl{
		authRepo: authRepo,
	}
}

func (s *authenticationServiceImpl) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {

	return &auth.LoginResponse{}, nil
}

func (s *authenticationServiceImpl) Logout(ctx context.Context, req *auth.LogoutRequest) (*auth.LogoutResponse, error) {
	// Implement logout logic
	return &auth.LogoutResponse{}, nil
}

func (s *authenticationServiceImpl) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	// Implement registration logic
	return &auth.RegisterResponse{}, nil
}

func (s *authenticationServiceImpl) ResetPassword(ctx context.Context, req *auth.ResetPasswordRequest) (*auth.ResetPasswordResponse, error) {
	// Implement password reset logic
	return &auth.ResetPasswordResponse{}, nil
}

func (s *authenticationServiceImpl) ChangePassword(ctx context.Context, req *auth.ChangePasswordRequest) (*auth.ChangePasswordResponse, error) {
	// Implement password change logic
	return &auth.ChangePasswordResponse{}, nil
}

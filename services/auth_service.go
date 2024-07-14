package services

import (
	auth "authentication-service/genproto/authentication_service"
	"authentication-service/models"
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
	authRepo     postgres.AuthenticationRepository
	tokenService TokenService
}

func NewAuthenticationService(userRepo postgres.UserRepository, authRepo postgres.AuthenticationRepository, tokenService TokenService) AuthenticationService {
	return &authenticationServiceImpl{
		authRepo:     authRepo,
		tokenService: tokenService,
	}
}

func (s *authenticationServiceImpl) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	var userModler models.User

	userModler.FullName = req.GetFullName()
	userModler.Username = req.GetUsername()
	userModler.Email = req.GetEmail()
	userModler.Bio = req.GetBio()
	userModler.UserType = req.GetUserType()
	userModler.PasswordHash = req.GetPassword() // original password

	err := s.authRepo.Register(ctx, &userModler)
	if err != nil {
		return nil, err
	}

	return &auth.RegisterResponse{}, nil
}

func (s *authenticationServiceImpl) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	user, err := s.authRepo.Login(ctx, req.Username, req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	var gReqToken auth.GenerateTokenRequest
	gReqToken.Email = user.Email
	gReqToken.Username = user.Username
	gReqToken.Password = user.PasswordHash

	resToken, err := s.tokenService.GenerateToken(ctx, &gReqToken)
	if err != nil {
		return nil, err
	}

	return &auth.LoginResponse{
		AccessToken:  resToken.AccessToken,
		RefreshToken: resToken.RefreshToken,
	}, nil
}

func (s *authenticationServiceImpl) Logout(ctx context.Context, req *auth.LogoutRequest) (*auth.LogoutResponse, error) {
	// Implement logout logic
	return &auth.LogoutResponse{}, nil
}
func (s *authenticationServiceImpl) ResetPassword(ctx context.Context, req *auth.ResetPasswordRequest) (*auth.ResetPasswordResponse, error) {
	// Implement password reset logic
	return &auth.ResetPasswordResponse{}, nil
}

func (s *authenticationServiceImpl) ChangePassword(ctx context.Context, req *auth.ChangePasswordRequest) (*auth.ChangePasswordResponse, error) {
	// Implement password change logic
	return &auth.ChangePasswordResponse{}, nil
}

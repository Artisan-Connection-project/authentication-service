package services

import (
	auth "authentication-service/genproto/authentication_service"
	"authentication-service/models"
	"authentication-service/storage/postgres"
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type AuthenticationService interface {
	Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error)
	Logout(ctx context.Context, req *auth.LogoutRequest) (*auth.LogoutResponse, error)
	Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error)
	ResetPassword(ctx context.Context, req *auth.ResetPasswordRequest) (*auth.ResetPasswordResponse, error)
	ChangePassword(ctx context.Context, req *auth.ChangePasswordRequest) (*auth.ChangePasswordResponse, error)
	VerifyEmailCodeAndUpdateUserInfo(ctx context.Context, email string) error
}

type authenticationServiceImpl struct {
	auth.UnimplementedAuthenticationServiceServer
	authRepo     postgres.AuthenticationRepository
	tokenService TokenService
	redisClient  *redis.Client
	emailService EmailService
}

func NewAuthenticationService(authRepo postgres.AuthenticationRepository, tokenService TokenService, emailService EmailService) AuthenticationService {
	return &authenticationServiceImpl{
		authRepo:     authRepo,
		tokenService: tokenService,

		emailService: emailService,
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
	gReqToken.Eamil = user.Email
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
	return &auth.LogoutResponse{}, nil
}

func (s *authenticationServiceImpl) ChangePassword(ctx context.Context, req *auth.ChangePasswordRequest) (*auth.ChangePasswordResponse, error) {
	return &auth.ChangePasswordResponse{}, nil
}

func (s *authenticationServiceImpl) ResetPassword(ctx context.Context, req *auth.ResetPasswordRequest) (*auth.ResetPasswordResponse, error) {
	user, err := s.authRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	resetCode := generateRandomCode()

	err = s.redisClient.Set(fmt.Sprintf("reset_code:%s", user.Email), resetCode, time.Minute*10).Err()
	if err != nil {
		return nil, fmt.Errorf("failed to store reset code")
	}

	err = s.emailService.SendCode(user.Email, resetCode)
	if err != nil {
		return nil, fmt.Errorf("failed to send reset code")
	}

	return &auth.ResetPasswordResponse{Message: "Reset code sent to email"}, nil
}

func (s *authenticationServiceImpl) VerifyResetCode(ctx context.Context, req *auth.VerifyResetCodeRequest) (*auth.VerifyResetCodeResponse, error) {
	storedCode, err := s.redisClient.Get(fmt.Sprintf("reset_code:%s", req.Email)).Result()
	if err != nil {
		return nil, fmt.Errorf("invalid or expired reset code")
	}

	if req.Code != storedCode {
		return nil, fmt.Errorf("invalid reset code")
	}

	err = s.authRepo.UpdatePassword(ctx, req.Email, req.NewPassword)
	if err != nil {
		return nil, fmt.Errorf("failed to update password")
	}

	err = s.redisClient.Del(fmt.Sprintf("reset_code:%s", req.Email)).Err()
	if err != nil {
		return nil, fmt.Errorf("failed to delete reset code")
	}

	return &auth.VerifyResetCodeResponse{Message: "Password reset successfully"}, nil
}

func generateRandomCode() string {
	return "123456"
}

func (s *authenticationServiceImpl) VerifyEmailCodeAndUpdateUserInfo(ctx context.Context, email string) error {
	return s.authRepo.UpdateUserToActive(ctx, email)
}

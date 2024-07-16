package services

import (
	auth "authentication-service/genproto/authentication_service"
	"context"
)

type MainService interface {
	ChangePassword(context.Context, *auth.ChangePasswordRequest) (*auth.ChangePasswordResponse, error)
	DeleteUserInfo(context.Context, *auth.DeleteUserInfoRequest) (*auth.DeleteUserInfoResponse, error)
	GenerateToken(context.Context, *auth.GenerateTokenRequest) (*auth.GenerateTokenResponse, error)
	GetUserInfo(context.Context, *auth.GetUserInfoRequest) (*auth.GetUserInfoResponse, error)
	GetUsersInfo(context.Context, *auth.GetUsersInfoRequest) (*auth.GetUsersInfoResponse, error)
	Login(context.Context, *auth.LoginRequest) (*auth.LoginResponse, error)
	Logout(context.Context, *auth.LogoutRequest) (*auth.LogoutResponse, error)
	RefreshToken(context.Context, *auth.RefreshTokenRequest) (*auth.RefreshTokenResponse, error)
	Register(context.Context, *auth.RegisterRequest) (*auth.RegisterResponse, error)
	ResetPassword(context.Context, *auth.ResetPasswordRequest) (*auth.ResetPasswordResponse, error)
	UpdateUserInfo(context.Context, *auth.UpdateUserInfoRequest) (*auth.UpdateUserInfoResponse, error)
	VerifyToken(context.Context, *auth.VerifyTokenRequest) (*auth.VerifyTokenResponse, error)
}
type mainServiceImpl struct {
	auth.UnimplementedAuthenticationServiceServer
	authService   AuthenticationService
	userService   UserManagementService
	tokenServices TokenService
}

func NewMainService(
	tokenServices TokenService,
	authService AuthenticationService,
	userService UserManagementService) MainService {
	return &mainServiceImpl{
		authService:   authService,
		userService:   userService,
		tokenServices: tokenServices,
	}
}

func (m *mainServiceImpl) ChangePassword(c context.Context, req *auth.ChangePasswordRequest) (*auth.ChangePasswordResponse, error) {
	return m.authService.ChangePassword(c, req)
}

func (m *mainServiceImpl) DeleteUserInfo(c context.Context, req *auth.DeleteUserInfoRequest) (*auth.DeleteUserInfoResponse, error) {
	return m.userService.DeleteUserInfo(c, req)
}

func (m *mainServiceImpl) GenerateToken(c context.Context, req *auth.GenerateTokenRequest) (*auth.GenerateTokenResponse, error) {
	return m.tokenServices.GenerateToken(c, req)
}
func (m *mainServiceImpl) GetUserInfo(c context.Context, req *auth.GetUserInfoRequest) (*auth.GetUserInfoResponse, error) {
	return m.userService.GetUserInfo(c, req)
}
func (m *mainServiceImpl) GetUsersInfo(c context.Context, req *auth.GetUsersInfoRequest) (*auth.GetUsersInfoResponse, error) {
	return m.userService.GetUsersInfo(c, req)
}

func (m *mainServiceImpl) Login(c context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	return m.authService.Login(c, req)
}

func (m *mainServiceImpl) Logout(c context.Context, req *auth.LogoutRequest) (*auth.LogoutResponse, error) {
	return m.authService.Logout(c, req)
}
func (m *mainServiceImpl) RefreshToken(c context.Context, req *auth.RefreshTokenRequest) (*auth.RefreshTokenResponse, error) {
	return m.tokenServices.RefreshToken(c, req.UserId, req)
}
func (m *mainServiceImpl) Register(c context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	return m.authService.Register(c, req)
}
func (m *mainServiceImpl) ResetPassword(c context.Context, req *auth.ResetPasswordRequest) (*auth.ResetPasswordResponse, error) {
	return m.authService.ResetPassword(c, req)
}
func (m *mainServiceImpl) UpdateUserInfo(c context.Context, req *auth.UpdateUserInfoRequest) (*auth.UpdateUserInfoResponse, error) {
	return m.userService.UpdateUserInfo(c, req)
}

func (m *mainServiceImpl) VerifyToken(c context.Context, req *auth.VerifyTokenRequest) (*auth.VerifyTokenResponse, error) {
	return m.tokenServices.VerifyToken(c, "", req)
}

package services

import (
	auth "authentication-service/genproto/authentication_service"
	"authentication-service/models"
	"authentication-service/storage/postgres"
	"context"
	"fmt"
	"time"

	"errors"

	"strconv"
)

type UserManagementService interface {
	GetUsersInfo(ctx context.Context, req *auth.GetUsersInfoRequest) (*auth.GetUsersInfoResponse, error)
	UpdateUserInfo(ctx context.Context, req *auth.UpdateUserInfoRequest) (*auth.UpdateUserInfoResponse, error)
	GetUserInfo(ctx context.Context, req *auth.GetUserInfoRequest) (*auth.GetUserInfoResponse, error)
	DeleteUserInfo(ctx context.Context, req *auth.DeleteUserInfoRequest) (*auth.DeleteUserInfoResponse, error)
}

type userManagementServiceImpl struct {
	userRepo postgres.UserRepository
	auth.UnimplementedAuthenticationServiceServer
}

func NewUserManagementService(userRepo postgres.UserRepository) UserManagementService {
	return &userManagementServiceImpl{
		userRepo: userRepo,
	}
}

func (s *userManagementServiceImpl) GetUsersInfo(ctx context.Context, req *auth.GetUsersInfoRequest) (*auth.GetUsersInfoResponse, error) {
	ctxCancel, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	page, err := strconv.Atoi(req.Page)
	if err != nil {
		return nil, fmt.Errorf("invalid page number: %w", err)
	}

	limit, err := strconv.Atoi(req.Limit)
	if err != nil {
		return nil, fmt.Errorf("invalid limit number: %w", err)
	}

	orderBy := req.OrderBy
	if orderBy == "" {
		orderBy = "id"
	}

	if page <= 0 || limit <= 0 {
		return nil, errors.New("page and limit must be greater than zero")
	}

	users, err := s.userRepo.GetUsersInfo(ctxCancel, page, limit, orderBy)
	if err != nil {
		return nil, fmt.Errorf("failed to get users info: %w", err)
	}

	return &auth.GetUsersInfoResponse{Users: users}, nil
}

func (s *userManagementServiceImpl) UpdateUserInfo(ctx context.Context, req *auth.UpdateUserInfoRequest) (*auth.UpdateUserInfoResponse, error) {
	user := &models.User{
		ID:       req.Id,
		Username: req.Username,
		FullName: req.FullName,
		Bio:      req.Bio,
		UserType: req.UserType,
		Email:    req.Email,
	}

	res, err := s.userRepo.UpdateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *userManagementServiceImpl) GetUserInfo(ctx context.Context, req *auth.GetUserInfoRequest) (*auth.GetUserInfoResponse, error) {
	user, err := s.userRepo.GetUserByID(ctx, req.Id)
	if err != nil {

		return nil, err
	}
	return &auth.GetUserInfoResponse{
		User: &auth.User{
			Id:        user.ID,
			Username:  user.Username,
			FullName:  user.FullName,
			Bio:       user.Bio,
			UserType:  user.UserType,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	}, nil
}

func (s *userManagementServiceImpl) DeleteUserInfo(ctx context.Context, req *auth.DeleteUserInfoRequest) (*auth.DeleteUserInfoResponse, error) {
	err := s.userRepo.DeleteUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &auth.DeleteUserInfoResponse{}, nil
}

package grpcservice

import (
	"context"

	"github.com/nikola-enter21/wms/backend/auth"

	userv1 "github.com/nikola-enter21/wms/backend/api/user/v1"
	"github.com/nikola-enter21/wms/backend/convert"
)

func (s *Server) CreateUser(ctx context.Context, in *userv1.CreateUserRequest) (*userv1.CreateUserResponse, error) {
	candidate, err := convert.CreateUserRequestToModel(in)
	if err != nil {
		return nil, err
	}

	inserted, err := s.DB.CreateUser(ctx, candidate)
	if err != nil {
		return nil, err
	}

	proto, err := convert.UserModelToProto(inserted)
	if err != nil {
		return nil, err
	}

	return &userv1.CreateUserResponse{
		User: proto,
	}, nil
}

func (s *Server) LoginUser(ctx context.Context, in *userv1.LoginUserRequest) (*userv1.LoginUserResponse, error) {
	fetched, err := s.DB.GetUserByCredentials(ctx, in.GetUsername(), in.GetPassword())
	if err != nil {
		return nil, err
	}

	accessToken, refreshToken, err := auth.GenerateTokens(fetched.ID, string(fetched.Role))
	if err != nil {
		return nil, err
	}

	proto, err := convert.UserModelToProto(fetched)
	if err != nil {
		return nil, err
	}

	return &userv1.LoginUserResponse{
		User:         proto,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *Server) ListUsers(ctx context.Context, in *userv1.ListUsersRequest) (*userv1.ListUsersResponse, error) {
	users, err := s.DB.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	protoUsers, err := convert.UsersModelToProto(users)
	if err != nil {
		return nil, err
	}

	return &userv1.ListUsersResponse{
		User: protoUsers,
	}, nil
}

func (s *Server) DeleteUser(ctx context.Context, in *userv1.DeleteUserRequest) (*userv1.DeleteUserResponse, error) {
	err := s.DB.DeleteUser(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &userv1.DeleteUserResponse{}, nil
}

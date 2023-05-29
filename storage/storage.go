package storage

import (
	"context"
	"practice1/user_service_go/genproto/user_service"
)

type StorageI interface {
	CloseDB()
	User() UserRepoI
}

type UserRepoI interface {
	Create(ctx context.Context, req *user_service.CreateUserRequest) (resp *user_service.Pkey, err error)
	GetById(ctx context.Context, req *user_service.Pkey) (resp *user_service.User, err error)
	GetAll(ctx context.Context, req *user_service.GetAllUsersRequest) (resp *user_service.GetAllUsersResponse, err error)
	Update(ctx context.Context, req *user_service.UpdateUserRequest) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *user_service.Pkey) (rowsAffected int64, err error)
}

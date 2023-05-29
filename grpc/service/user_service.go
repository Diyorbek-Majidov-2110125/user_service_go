package service

import (
	"context"
	"fmt"
	"practice1/user_service_go/config"
	"practice1/user_service_go/genproto/user_service"
	"practice1/user_service_go/grpc/client"
	"practice1/user_service_go/pkg/logger"
	"practice1/user_service_go/storage"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type userService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	user_service.UnimplementedUserServiceServer
}

func NewUserService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, svcs client.ServiceManagerI) *userService {
	return &userService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: svcs,
	}
}

func (b *userService) Create(ctx context.Context, req *user_service.CreateUserRequest) (resp *user_service.User, err error) {
	b.log.Info("---CreateUser--->", logger.Any("req", req))

	pKey, err := b.strg.User().Create(ctx, req)

	if err != nil {
		b.log.Error("!!!CreateUser--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	fmt.Println("why")
	resp = &user_service.User{
		Id:          pKey.Id,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		PhoneNumber: req.PhoneNumber,
	}
	return resp, nil
}

func (b *userService) GetById(ctx context.Context, req *user_service.Pkey) (resp *user_service.User, err error) {
	b.log.Info("---GetUserById--->", logger.Any("req", req))

	resp, err = b.strg.User().GetById(ctx, req)

	if err != nil {
		b.log.Error("!!!GetUserById--->", logger.Error(err))
		return resp, status.Error(codes.InvalidArgument, err.Error())
	}
	return resp, nil
}

func (b *userService) GetAll(ctx context.Context, req *user_service.GetAllUsersRequest) (resp *user_service.GetAllUsersResponse, err error) {
	b.log.Info("---GetAllUsers--->", logger.Any("req", req))

	resp, err = b.strg.User().GetAll(ctx, req)

	if err != nil {
		b.log.Error("!!!GetAllUsers--->", logger.Error(err))
		return resp, status.Error(codes.InvalidArgument, err.Error())
	}
	return resp, nil
}
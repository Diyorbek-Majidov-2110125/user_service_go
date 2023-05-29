package grpc

import (
	"practice1/user_service_go/config"
	"practice1/user_service_go/genproto/user_service"
	"practice1/user_service_go/grpc/client"
	"practice1/user_service_go/grpc/service"
	"practice1/user_service_go/pkg/logger"
	"practice1/user_service_go/storage"

	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI, svcs client.ServiceManagerI) (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer()

	user_service.RegisterUserServiceServer(grpcServer, service.NewUserService(cfg, log, strg, svcs))

	reflection.Register(grpcServer)
	return
}

package client

import "practice1/user_service_go/config"

type ServiceManagerI interface {
	// OrderService() order_service.OrderServiceClient
}

type grpcClients struct {
	// orderService order_service.OrderServiceClient
}

func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {
	// connOrderService, err := grpc.Dial(
	// 	cfg.ServiceHost+cfg.ServicePort,
	// 	grpc.WithInsecure(),
	// )
	// if err != nil {
	// 	return nil, err
	// }

	// if err != nil {
	// 	return nil, err
	// }
	return &grpcClients{
		// orderService: order_service.NewOrderServiceClient(connOrderService),
	}, nil
}

// func (c *grpcClients) OrderService() order_service.OrderServiceClient {
// 	return c.orderService
// }

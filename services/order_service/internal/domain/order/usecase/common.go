package order_usecase

import (
	"context"
	"github.com/serenite11/market/proto/services/order_service_v1"
	"github.com/serenite11/market/services/order-service/internal/clients"
	"github.com/serenite11/market/services/order-service/internal/factory"
)

type uc struct {
	factory *factory.Factory
	clients *clients.GClients
}

var _ UseCase = (*uc)(nil)

func New(factory *factory.Factory, gClients *clients.GClients) *uc {
	return &uc{
		factory: factory,
		clients: gClients,
	}
}

type (
	UseCase interface {
		CreateOrder(ctx context.Context, request *order_service_v1.CreateOrder_Request) (*order_service_v1.CreateOrder_Response, error)
		GetOrderById(ctx context.Context, request *order_service_v1.GetOrderById_Request) (*order_service_v1.GetOrderById_Response, error)
		FetchOrders(ctx context.Context) error
	}
)

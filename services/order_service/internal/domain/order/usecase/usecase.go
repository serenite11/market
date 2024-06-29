package usecase

import (
	"context"
	"github.com/serenite11/market/proto/services/order_service_v1"
	"order-service/internal/factory"
)

type uc struct {
	factory *factory.Factory
}

var _ UseCase = (*uc)(nil)

func New(factory *factory.Factory) *uc {
	return &uc{
		factory: factory,
	}
}

func (u uc) CreateOrder(ctx context.Context, request order_service_v1.CreateOrder_Request) (string, error) {

	u.factory.OrderRepo().CreateOrder(ctx)
}

func (u uc) GetOrderById() (order_service_v1.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (u uc) FetchOrders() error {
	//TODO implement me
	panic("implement me")
}

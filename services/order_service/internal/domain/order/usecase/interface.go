package usecase

import (
	"context"
	"github.com/serenite11/market/proto/services/order_service_v1"
)

type (
	UseCase interface {
		CreateOrder(ctx context.Context, request order_service_v1.CreateOrder_Request) (string, error)
		GetOrderById() (order_service_v1.Order, error)
		FetchOrders() error
	}
)

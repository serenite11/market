package order

import (
	"context"
	"github.com/serenite11/market/proto/services/order-service-v1"
)

type (
	UseCase interface {
		CreateOrder(
			ctx context.Context,
			request *order_service_v1.CreateOrder_Request,
		) (*order_service_v1.CreateOrder_Response, error)
		GetOrderById(
			ctx context.Context,
			request *order_service_v1.GetOrderById_Request,
		) (*order_service_v1.GetOrderById_Response, error)
		FetchOrders(
			ctx context.Context,
			request *order_service_v1.FetchOrdersByUserId_Request,
		) (*order_service_v1.FetchOrdersByUserId_Response, error)
	}
)

package usecase

import "github.com/serenite11/market/proto/services/order_service_v1"

type (
	Order interface {
		CreateOrder(request order_service_v1.CreateOrder_Request) (string, error)
		GetOrderById() (order_service_v1.Order, error)
		FetchOrders() error
	}
)

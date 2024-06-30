package delivery_grpc

import (
	"context"
	"github.com/serenite11/market/proto/services/order_service_v1"
	order_usecase "order-service/internal/domain/order/usecase"
)

type handler struct {
	uc order_usecase.UseCase
	order_service_v1.UnimplementedOrderServiceServer
}

func NewOrderHandler(orderUc order_usecase.UseCase) order_service_v1.OrderServiceServer {
	return &handler{
		uc: orderUc,
	}
}

func (h handler) CreateOrder(ctx context.Context, request *order_service_v1.CreateOrder_Request) (*order_service_v1.CreateOrder_Response, error) {
	return h.uc.CreateOrder(ctx, request)
}
func (h handler) GetOrderById(ctx context.Context, request *order_service_v1.GetOrderById_Request) (*order_service_v1.GetOrderById_Response, error) {
	return h.uc.GetOrderById(ctx, request)
}

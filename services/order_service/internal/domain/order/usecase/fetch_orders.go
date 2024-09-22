package order_usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/serenite11/market/proto/services/order_service_v1"
)

func (u uc) FetchOrders(
	ctx context.Context,
	request *order_service_v1.FetchOrdersByUserId_Request,
) (*order_service_v1.FetchOrdersByUserId_Response, error) {
	userId, err := uuid.Parse(request.UserId)
	if err != nil {
		return nil, err
	}
	orders, err := u.factory.OrderRepo().FetchByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}
	return &order_service_v1.FetchOrdersByUserId_Response{
		Orders: u.Mapper().FromEntityOrdersToProto(orders),
	}, nil
}

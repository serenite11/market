package order_usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/serenite11/market/proto/services/order_service_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (u uc) FetchOrders(ctx context.Context, request *order_service_v1.FetchOrdersByUserId_Request) (*order_service_v1.FetchOrdersByUserId_Response, error) {
	userId, err := uuid.Parse(request.UserId)
	if err != nil {
		return nil, err
	}
	orders, err := u.factory.OrderRepo().FetchByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}
	respOrders := make([]*order_service_v1.Order, len(orders))		
	for _,order:=range orders{
			respOrders = append(
				respOrders,
				&order_service_v1.Order{
					Id:          order.Id.String(),
					Amount:      order.Amount,
					Status:      order.Status,
					CreatedAt:   timestamppb.New(order.CreatedAt),
					UpdatedAt:   timestamppb.New(order.UpdatedAt),
					CompletedAt: timestamppb.New(order.CompletedAt.Time),
				},
			)
	}

	return &order_service_v1.FetchOrdersByUserId_Response{
		Orders: []*order_service_v1.Order{},
	},nil
}

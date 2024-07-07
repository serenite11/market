package order_usecase

import (
	"context"
	"github.com/google/uuid"
	"github.com/serenite11/market/proto/services/order_service_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (u uc) GetOrderById(ctx context.Context, request *order_service_v1.GetOrderById_Request) (*order_service_v1.GetOrderById_Response, error) {
	id, err := uuid.Parse(request.GetId())
	if err != nil {
		return nil, err
	}

	order, err := u.factory.OrderRepo().GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &order_service_v1.GetOrderById_Response{Order: &order_service_v1.Order{
		Id:          order.Id.String(),
		Amount:      order.Amount,
		Status:      order.Status,
		CreatedAt:   timestamppb.New(order.CreatedAt),
		UpdatedAt:   timestamppb.New(order.UpdatedAt),
		CompletedAt: timestamppb.New(order.CompletedAt.Time),
	}}, nil
}

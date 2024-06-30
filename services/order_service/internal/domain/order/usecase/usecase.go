package order_usecase

import (
	"context"
	"github.com/google/uuid"
	"github.com/serenite11/market/proto/services/order_service_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"order-service/internal/clients"
	order_model "order-service/internal/domain/order/model"
	"order-service/internal/factory"
	"time"
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

func (u uc) CreateOrder(ctx context.Context, request *order_service_v1.CreateOrder_Request) (*order_service_v1.CreateOrder_Response, error) {
	userId, err := uuid.Parse(request.UserId)
	if err != nil {
		return nil, ErrInvalidUser
	}

	var amount float64

	for _, item := range request.GetProducts() {
		amount += item.Price * float64(item.Quantity)
	}

	order := order_model.Order{}

	order.
		SetId(uuid.Must(uuid.NewV7())).
		SetUserId(userId).
		SetAmount(amount)

	ctxCreateOrder, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()
	orderId, err := u.factory.OrderRepo().Create(ctxCreateOrder, &order)
	if err != nil {
		return nil, err
	}

	return &order_service_v1.CreateOrder_Response{
		OrderId: orderId.String(),
		Status:  order_service_v1.OrderStatus_CREATED,
	}, nil
}

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

func (u uc) FetchOrders(ctx context.Context) error {
	return nil
}

package usecase

import (
	"context"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/serenite11/market/proto/services/order_service_v1"
	"github.com/serenite11/market/proto/services/product_storage_v1"
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

func (u uc) CreateOrder(ctx context.Context, request order_service_v1.CreateOrder_Request) (*order_service_v1.CreateOrder_Response, error) {

	for _, item := range request.Products {
		product, err := u.clients.ProductStorageClient().GetProductById(ctx, &product_storage_v1.GetProductById_Request{
			Id: item.Id,
		})
		if err != nil {
			return nil, err
		}
	}

	u.factory.OrderRepo().Create(ctx, &order_model.Order{
		Id:          uuid.Must(uuid.NewV7()),
		Amount:      0,
		UserId:      uuid.UUID{},
		Products:    nil,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		CompletedAt: pq.NullTime{},
	})

	return &order_service_v1.CreateOrder_Response{
		OrderId: "",
		Status:  0,
	}, nil
}

func (u uc) GetOrderById(ctx context.Context) (order_service_v1.Order, error) {
	u.factory.OrderRepo().GetById(ctx)
}

func (u uc) FetchOrders() error {
	//TODO implement me
	panic("implement me")
}

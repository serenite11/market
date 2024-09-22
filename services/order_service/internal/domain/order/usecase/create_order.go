package order_usecase

import (
	"context"
	"github.com/serenite11/market/services/order-service/internal/infra/factory"
	"time"

	"github.com/google/uuid"
	"github.com/serenite11/market/proto/services/order_service_v1"
	order_model "github.com/serenite11/market/services/order-service/internal/domain/order/model"
)

func (u uc) CreateOrder(
	ctx context.Context,
	request *order_service_v1.CreateOrder_Request,
) (*order_service_v1.CreateOrder_Response, error) {
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
		SetAmount(amount).
		SetProducts(request.GetProducts())

	ctxCreateOrder, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	orderId := uuid.UUID{}

	if err = u.factory.RunTx(ctx, func(tx *factory.Factory) error {
		oId, err := tx.OrderRepo().Create(ctxCreateOrder, &order)
		if err != nil {
			return err
		}
		orderId = *oId
		return nil
	}); err != nil {
		return nil, err
	}

	return &order_service_v1.CreateOrder_Response{
		OrderId: orderId.String(),
		Status:  order_service_v1.OrderStatus_CREATED,
	}, nil
}

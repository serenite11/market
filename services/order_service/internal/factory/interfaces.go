package factory

import (
	"context"
	"github.com/google/uuid"
	order_model "order-service/internal/domain/order/model"
)

type (
	OrderRepository interface {
		CreateOrder(ctx context.Context, order *order_model.Order) (*uuid.UUID, error)
		GetById(ctx context.Context, id uuid.UUID) (*order_model.Order, error)
	}
)

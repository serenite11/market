package factory

import (
	"context"
	"github.com/google/uuid"
	order_model "order-service/internal/domain/order/model"
)

//go:generate mockgen -source interfaces.go ReposMocks  -destination ./mocks
type (
	OrderRepository interface {
		Create(ctx context.Context, order *order_model.Order) (*uuid.UUID, error)
		GetById(ctx context.Context, id uuid.UUID) (*order_model.Order, error)
		FetchByUserId(ctx context.Context, userId uuid.UUID) ([]*order_model.Order, error)
	}
)

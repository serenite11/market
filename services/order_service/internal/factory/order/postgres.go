package order_repository

import (
	"context"
	"github.com/google/uuid"
	order_model "order-service/internal/domain/order/model"
	storage_postgres "order-service/pkg/storage/postgres"
)

type postgres struct {
	db storage_postgres.Postgres
}

func InitPostgres(db storage_postgres.Postgres) *postgres {
	return &postgres{db: db}
}

const createOrderQuery = `INSERT INTO order (id,amount) values ($1,$2) RETURNING id`

func (p postgres) CreateOrder(ctx context.Context, order *order_model.Order) (*uuid.UUID, error) {
	var data uuid.UUID

	if err := p.db.QueryRow(ctx, createOrderQuery, order.Id, order.Amount).Scan(&data); err != nil {
		return nil, err
	}

	return &data, nil
}

const getByIdQuery = `SELECT id,amount,created_at,updated_at FROM "order" WHERE id = $1;`

func (p postgres) GetById(ctx context.Context, id uuid.UUID) (*order_model.Order, error) {
	var data order_model.Order

	if err := p.db.Get(ctx, &data, getByIdQuery, id); err != nil {
		return nil, err
	}
	return &data, nil
}

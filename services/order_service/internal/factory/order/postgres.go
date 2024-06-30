package order_repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/huandu/go-sqlbuilder"
	order_model "order-service/internal/domain/order/model"
	storage_postgres "order-service/pkg/storage/postgres"
)

type postgres struct {
	db storage_postgres.Postgres
}

func InitPostgres(db storage_postgres.Postgres) *postgres {
	return &postgres{db: db}
}

const createOrderQuery = `INSERT INTO order (id,amount,created_at,updated_at) values ($1,$2,now(),now()) RETURNING id`

func (p postgres) Create(ctx context.Context, order *order_model.Order) (*uuid.UUID, error) {
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

func (p postgres) Update(ctx context.Context) error {
	queryBuild := sqlbuilder.PostgreSQL.NewUpdateBuilder().Update(`"order"`)

	query, args := queryBuild.Build()

	_, err := p.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

const fetchByUserIdQuery = `SELECT id,amount,created_at FROM "order" WHERE id = $1;`

func (p postgres) FetchByUserId(ctx context.Context, userId uuid.UUID) ([]*order_model.Order, error) {
	var orders []*order_model.Order

	err := p.db.Select(ctx, &orders, fetchByUserIdQuery, userId)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

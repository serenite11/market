package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/huandu/go-sqlbuilder"
	order_model "github.com/serenite11/market/services/order-service/internal/domain/order/model"
	storage_postgres "github.com/serenite11/market/services/order-service/pkg/storage/postgres"
)

type orderRepo struct {
	db storage_postgres.Postgres
}

func InitOrderRepo(db storage_postgres.Postgres) *orderRepo {
	return &orderRepo{db: db}
}

const createOrderQuery = `INSERT INTO order (id,amount,created_at,updated_at) values ($1,$2,now(),now()) RETURNING id`

func (p orderRepo) Create(ctx context.Context, order *order_model.Order) (*uuid.UUID, error) {
	var data uuid.UUID

	if err := p.db.QueryRow(ctx, createOrderQuery, order.Id, order.Amount).Scan(&data); err != nil {
		return nil, err
	}

	return &data, nil
}

const getByIdQuery = `SELECT id,amount,created_at,updated_at FROM "order" WHERE id = $1;`

func (p orderRepo) GetById(ctx context.Context, id uuid.UUID) (*order_model.Order, error) {
	var data order_model.Order

	if err := p.db.Get(ctx, &data, getByIdQuery, id); err != nil {
		return nil, err
	}
	return &data, nil
}

func (p orderRepo) Update(ctx context.Context) error {
	queryBuild := sqlbuilder.PostgreSQL.NewUpdateBuilder().Update(`"order"`)

	query, args := queryBuild.Build()

	_, err := p.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

const fetchByUserIdQuery = `SELECT id,amount,created_at FROM "order" WHERE user_id = $1;`

func (p orderRepo) FetchByUserId(
	ctx context.Context,
	userId uuid.UUID,
) ([]*order_model.Order, error) {
	var orders []*order_model.Order

	err := p.db.Select(ctx, &orders, fetchByUserIdQuery, userId)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

const createOutboxQuery = `INSERT INTO outbox (created_at,routing_key,message) VALUES (now(), $1, $2)`

func (p orderRepo) CreateOutbox(ctx context.Context, routingKey string, msg []byte) error {
	_, err := p.db.Exec(ctx, createOutboxQuery, routingKey, msg)
	if err != nil {
		return err
	}

	return nil
}

const fetchoutboxmessagesQuery = `SELECT id,`

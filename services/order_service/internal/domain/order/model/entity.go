package order_model

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/serenite11/market/proto/services/order_service_v1"
	"time"
)

type Order struct {
	Id          uuid.UUID                    `db:"id"`
	Amount      float64                      `db:"amount"`
	UserId      uuid.UUID                    `db:"user_id"`
	Status      order_service_v1.OrderStatus `db:"status"`
	Products    any                          `db:"products"`
	CreatedAt   time.Time                    `db:"created_at"`
	UpdatedAt   time.Time                    `db:"updated_at"`
	CompletedAt pq.NullTime                  `db:"completed_at"`
}

func (o *Order) SetId(id uuid.UUID) *Order {
	o.Id = id
	return o
}

func (o *Order) SetAmount(amount float64) *Order {
	o.Amount = amount
	return o
}

func (o *Order) SetUserId(userId uuid.UUID) *Order {
	o.UserId = userId
	return o
}

func (o *Order) SetStatus(status order_service_v1.OrderStatus) *Order {
	o.Status = status
	return o
}

func (o *Order) SetProducts(products any) *Order {
	o.Products = products
	return o
}
func (o *Order) SetCreatedAt(createdAt time.Time) *Order {
	o.CreatedAt = createdAt
	return o
}
func (o *Order) SetUpdatedAt(updatedAt time.Time) *Order {
	o.UpdatedAt = updatedAt
	return o
}
func (o *Order) SetCompletedAt(completedAt time.Time) *Order {
	o.CompletedAt = pq.NullTime{}
	return o
}

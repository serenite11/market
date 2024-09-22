package order_model

import (
	"github.com/goccy/go-json"
	"time"

	"github.com/google/uuid"
	"github.com/guregu/null/v5"
	"github.com/serenite11/market/proto/services/order_service_v1"
)

type ProductsAl []*order_service_v1.ProductOrder

type Order struct {
	Id          uuid.UUID                    `db:"id"`
	Amount      float64                      `db:"amount"`
	UserId      uuid.UUID                    `db:"user_id"`
	Status      order_service_v1.OrderStatus `db:"status"`
	Products    null.Value[ProductsAl]       `db:"products"`
	CreatedAt   time.Time                    `db:"created_at"`
	UpdatedAt   time.Time                    `db:"updated_at"`
	CompletedAt null.Time                    `db:"completed_at"`
}

func (p *ProductsAl) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	return json.Unmarshal(value.([]byte), p)
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

func (o *Order) SetProducts(products []*order_service_v1.ProductOrder) *Order {
	o.Products = null.NewValue[ProductsAl](products, len(products) != 0)
	return o
}
func (o *Order) SetCreatedAt(createdAt time.Time) *Order {
	o.CreatedAt = createdAt
	return o
}
func (o *Order) SetCompleted() *Order {
	o.CompletedAt = null.TimeFrom(time.Now())
	return o
}

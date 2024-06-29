package order_model

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	"time"
)

type Order struct {
	Id          uuid.UUID   `db:"id"`
	Amount      float64     `db:"amount"`
	UserId      uuid.UUID   `db:"user_id"`
	Products    any         `db:"products"`
	CreatedAt   time.Time   `db:"created_at"`
	UpdatedAt   time.Time   `db:"updated_at"`
	CompletedAt pq.NullTime `db:"completed_at"`
}

package order_usecase

import (
	"context"
	"github.com/serenite11/market/services/order-service/internal/infrastructure/factory"
	"time"

	"github.com/serenite11/market/services/order-service/pkg/kafka"
)

type Outbox struct {
	prod    kafka.Producer
	period  time.Duration
	factory *factory.Factory
	done    chan struct{}
}

func NewOutbox(period time.Duration, producer kafka.Producer, f *factory.Factory) *Outbox {
	return &Outbox{
		prod:   producer,
		period: period,
		done:   make(chan struct{}),
	}
}

func (o *Outbox) Run(ctx context.Context) error {
	ticker := time.NewTicker(o.period)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-o.done:
				return
			case <-ticker.C:
				o.run(ctx)
			}
		}
	}()
	return nil
}

func (o *Outbox) run(ctx context.Context) {

}

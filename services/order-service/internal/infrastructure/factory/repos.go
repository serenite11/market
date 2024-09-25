package factory

import (
	"github.com/serenite11/market/services/order-service/internal/infrastructure/factory/postgres"
)

func (f *Factory) OrderRepo() OrderRepository {
	return postgres.InitOrderRepo(f.postgres)
}

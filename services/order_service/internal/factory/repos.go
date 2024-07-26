package factory

import order_repository "github.com/serenite11/market/services/order-service/internal/factory/order"

func (f *Factory) OrderRepo() OrderRepository {
	return order_repository.InitPostgres(f.postgres)
}

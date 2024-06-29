package factory

import order_repository "order-service/internal/factory/order"

func (f *Factory) OrderRepo() OrderRepository {
	return order_repository.InitPostgres(f.postgres)
}

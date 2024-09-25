package order_usecase

import (
	"github.com/serenite11/market/services/order-service/internal/domain/order/mapper"
	"github.com/serenite11/market/services/order-service/internal/infrastructure/factory"
	"github.com/serenite11/market/services/order-service/internal/infrastructure/grpc/clients"
)

type uc struct {
	factory *factory.Factory
	clients *clients.GClients
}

func New(factory *factory.Factory, gClients *clients.GClients) *uc {
	return &uc{
		factory: factory,
		clients: gClients,
	}
}

func (u uc) Mapper() mapper.Converter {
	return &mapper.OrderConverter{}
}

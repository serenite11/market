package order_usecase

import (
	"order-service/internal/clients"
	"order-service/internal/factory"
)

type uc struct {
	factory *factory.Factory
	clients *clients.GClients
}

var _ UseCase = (*uc)(nil)

func New(factory *factory.Factory, gClients *clients.GClients) *uc {
	return &uc{
		factory: factory,
		clients: gClients,
	}
}

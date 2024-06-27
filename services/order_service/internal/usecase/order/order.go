package order

import (
	"order-service/internal/factory"
)

type uc struct {
	factory *factory.Factory
}

func New(factory *factory.Factory) *uc {
	return &uc{
		factory: factory,
	}
}

func (u uc) CreateOrder() (*string, error) {

}

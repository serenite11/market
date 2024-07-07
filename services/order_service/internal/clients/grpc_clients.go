package clients

import "github.com/serenite11/market/services/order-service/internal/clients/product"

type GClients struct {
	productAPI *product.Client
}

func NewGClients() *GClients {
	return &GClients{}
}

func (g *GClients) ProductStorageClient() *product.Client {
	return g.productAPI
}

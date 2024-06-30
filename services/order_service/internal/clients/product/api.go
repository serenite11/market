package product

import (
	"context"
	"github.com/serenite11/market/proto/services/product_storage_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	api  product_storage_v1.ProductStorageClient
	conn *grpc.ClientConn
}

func NewClient(address string) (*Client, error) {
	var options []grpc.DialOption

	options = append(options, grpc.WithTransportCredentials(insecure.NewCredentials()))

	grpcConn, err := grpc.NewClient(address, options...)
	if err != nil {
		return nil, err
	}
	return &Client{
		api:  product_storage_v1.NewProductStorageClient(grpcConn),
		conn: grpcConn,
	}, nil
}

func (c Client) GetProductById(ctx context.Context, request *product_storage_v1.GetProductById_Request) (*product_storage_v1.GetProductById_Response, error) {
	return c.api.GetProductById(ctx, request)
}

func (c Client) FetchProducts(ctx context.Context, request *product_storage_v1.FetchProducts_Request) (*product_storage_v1.FetchProducts_Response, error) {
	return c.api.FetchProducts(ctx, request)
}

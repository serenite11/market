package orderservice

import (
	"context"
	orderservice_v1 "github.com/serenite11/market/proto/services/order-service-v1"
	"github.com/serenite11/market/services/api-gateway/config"
	"google.golang.org/grpc"
)

type Client struct {
	conn *grpc.ClientConn
	api  orderservice_v1.OrderServiceClient
	cfg  *config.OrderService
}

func NewClient(cfg *config.OrderService) *Client {
	return &Client{
		cfg: cfg,
	}
}

func (c *Client) OnStart(_ context.Context) error {
	var err error
	c.conn, err = grpc.NewClient(c.cfg.DSN)
	if err != nil {
		return err
	}
	c.api = orderservice_v1.NewOrderServiceClient(c.conn)
	return nil
}

func (c *Client) OnStop(_ context.Context) error {
	return c.conn.Close()
}

func (c *Client) CreateOrder(ctx context.Context, request *orderservice_v1.CreateOrder_Request) (*orderservice_v1.CreateOrder_Response, error) {
	return c.api.CreateOrder(ctx, request)
}
func (c *Client) GetOrderById(ctx context.Context, request *orderservice_v1.GetOrderById_Request) (*orderservice_v1.GetOrderById_Response, error) {
	return c.api.GetOrderById(ctx, request)
}
func (c *Client) FetchOrdersByUserId(ctx context.Context, request *orderservice_v1.FetchOrdersByUserId_Request) (*orderservice_v1.FetchOrdersByUserId_Response, error) {
	return &orderservice_v1.FetchOrdersByUserId_Response{}, nil
}

func (c *Client) CancelOrder(ctx context.Context, request *orderservice_v1.CancelOrder_Request) (*orderservice_v1.CancelOrder_Response, error) {
	return &orderservice_v1.CancelOrder_Response{}, nil
}

package orderservice

import (
	orderservice_v1 "github.com/serenite11/market/proto/services/order-service-v1"
	"google.golang.org/grpc"
)

type Client struct {
	conn *grpc.ClientConn
	api  orderservice_v1.OrderServiceClient
}

func NewClient()

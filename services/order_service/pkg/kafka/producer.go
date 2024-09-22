package kafka

import "context"

type Producer interface {
	Send(topic string, message []byte) error
	Connect(ctx context.Context) error
	Close(ctx context.Context) error
}

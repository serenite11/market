package kafka_producer

import (
	"context"
	"github.com/IBM/sarama"
	"go.uber.org/zap"
)

type Producer interface {
	Send(topic string, message []byte) error
	Connect(ctx context.Context) error
	Close(ctx context.Context) error
}

type producer struct {
	client sarama.SyncProducer
	log    *zap.Logger
}

func New(log *zap.Logger) *producer {
	return &producer{
		log: log,
	}
}

func (p *producer) Connect(_ context.Context) error {
	client, err := sarama.NewSyncProducer(
		[]string{"localhost:3000"},
		&sarama.Config{},
	)
	if err != nil {
		return err
	}
	p.client = client
	return nil
}

func (p *producer) Close(_ context.Context) error {
	return p.client.Close()
}

func (p *producer) Send(topic string, message []byte) error {
	partition, offset, err := p.client.SendMessage(&sarama.ProducerMessage{
		Topic:    topic,
		Value:    sarama.ByteEncoder(message),
		Headers:  nil,
		Metadata: nil,
	})
	if err != nil {
		return err
	}
	p.log.Info("send message to kafka",
		zap.Any("partition", partition),
		zap.Any("offset", offset),
		zap.String("topic", topic),
		zap.Any("message", string(message)),
	)

	return nil
}

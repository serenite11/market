package kafka

import kafka_consumer "github.com/serenite11/market/services/order-service/pkg/kafka/consumer"

type Server struct {
	c kafka_consumer.Consumer
}

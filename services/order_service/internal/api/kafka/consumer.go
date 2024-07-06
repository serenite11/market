package api_kafka

import kafka_consumer "order-service/pkg/kafka/consumer"

type Server struct {
	c kafka_consumer.Consumer
}

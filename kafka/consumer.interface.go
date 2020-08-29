package kafka

import "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"

type ConsumerInterface interface {
	Consume(msg *kafka.Message)
}

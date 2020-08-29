package kafka

import "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"

type ConsumerInterface interface {
	OnConsume(msg *kafka.Message)
}

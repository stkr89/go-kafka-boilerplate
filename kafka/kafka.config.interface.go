package kafka

import "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"

type KafkaConfigInterface interface {
	consume(topic string) (*kafka.Consumer, error)
}

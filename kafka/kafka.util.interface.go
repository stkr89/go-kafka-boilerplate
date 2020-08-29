package kafka

import "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"

type KafkaUtilInterface interface {
	consume(topic string) (*kafka.Consumer, error)
}

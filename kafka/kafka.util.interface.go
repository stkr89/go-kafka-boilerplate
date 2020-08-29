package kafka

import "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"

type KafkaUtilInterface interface {
	addConsumer(topic string, onConsume func(msg *kafka.Message))

	Produce(topic string, payload interface{})
}

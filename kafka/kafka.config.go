package kafka

import (
	"cloudpro-metadata-handler/common"
	"github.com/sirupsen/logrus"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"os"
)

type KafkaConfig struct {
}

func NewKafkaConfig() *KafkaConfig {
	return &KafkaConfig{}
}

func (k *KafkaConfig) consume(topic string, listener func(msg *kafka.Message)) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv(common.KafkaBootstrapServers),
		"group.id":          os.Getenv(common.KafkaGroupSSEProactiveSupport),
		"auto.offset.reset": common.KafkaAutoOffsetReset,
	})
	if err != nil {
		panic(err)
	}

	err = c.Subscribe(topic, nil)
	if err != nil {
		panic(err)
	}

	logrus.Infof("Subscribed to %s topic", topic)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			listener(msg)
		} else {
			logrus.Errorf("Consume error: %v (%v)\n", err, msg)
		}
	}

	c.Close()
}

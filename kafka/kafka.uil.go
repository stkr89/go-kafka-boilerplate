package kafka

import (
	"cloudpro-metadata-handler/common"
	"github.com/sirupsen/logrus"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"os"
)

type KafkaUtil struct {
}

func NewKafkaUtil() *KafkaUtil {
	return &KafkaUtil{}
}

func (ku *KafkaUtil) Produce(topic string, payload interface{}) {
	panic("implement me")
}

func (ku *KafkaUtil) addConsumer(topic string, onConsume func(msg *kafka.Message)) {
	c, err := kafka.NewConsumer(getBaseKafkaConfig())
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
			onConsume(msg)
		} else {
			logrus.Errorf("OnConsume error: %v (%v)\n", err, msg)
		}
	}

	c.Close()
}

func getBaseKafkaConfig() *kafka.ConfigMap {
	return &kafka.ConfigMap{
		"bootstrap.servers": os.Getenv(common.KafkaBootstrapServers),
		"group.id":          os.Getenv(common.KafkaTestGroup),
		"auto.offset.reset": common.KafkaAutoOffsetReset,
	}
}

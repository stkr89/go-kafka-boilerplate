package kafka

import (
	"cloudpro-metadata-handler/common"
	"cloudpro-metadata-handler/service/test"
)

type KafkaConsumer struct {
	kafkaUtil   *KafkaUtil
	testService *test.TestService
}

func NewKafkaConsumer() *KafkaConsumer {
	return &KafkaConsumer{
		kafkaUtil:   NewKafkaUtil(),
		testService: test.NewTestService(),
	}
}

func (kc *KafkaConsumer) InitConsumers() {
	listenersMap := map[string]ConsumerInterface{
		common.KafkaTopicTest: kc.testService,
	}

	for k, v := range listenersMap {
		kc.kafkaUtil.consume(k, v.Consume)
	}
}

package kafka

import (
	"cloudpro-metadata-handler/common"
	"cloudpro-metadata-handler/service/test"
)

type KafkaConsumer struct {
	kafkaConfig *KafkaConfig
	testService *test.TestService
}

func NewKafkaConsumer() *KafkaConsumer {
	return &KafkaConsumer{
		kafkaConfig: NewKafkaConfig(),
		testService: test.NewTestService(),
	}
}

func (kc *KafkaConsumer) InitConsumers() {
	listenersMap := map[string]ConsumerInterface{
		common.KafkaTopicTest: kc.testService,
	}

	for k, v := range listenersMap {
		kc.kafkaConfig.consume(k, v.Consume)
	}
}

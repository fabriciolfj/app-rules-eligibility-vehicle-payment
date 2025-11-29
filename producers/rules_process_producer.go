package producers

import (
	"github.com/IBM/sarama"
	"github.com/fabriciolfj/rules-elegibility/configuration"
)

type RuleProcessProducer struct {
	producer sarama.AsyncProducer
}

func ProviderRuleProcessProducer(cfg *configuration.KafkaConfig) (*RuleProcessProducer, error) {
	producer := &RuleProcessProducer{
		producer: cfg.Producer,
	}

	return producer, nil
}

func (producer *RuleProcessProducer) SendMessage(message string, topic string) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	producer.producer.Input() <- msg

	select {
	case err := <-producer.producer.Errors():
		return err
	case <-producer.producer.Successes():
		return nil
	}
}

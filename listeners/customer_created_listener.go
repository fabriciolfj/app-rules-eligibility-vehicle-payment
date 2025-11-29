package listeners

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/IBM/sarama"
	"github.com/fabriciolfj/rules-elegibility/configuration"
	"github.com/fabriciolfj/rules-elegibility/dtos"
	"github.com/magiconair/properties"
)

type CustomerCreatedListener struct {
	consumer sarama.ConsumerGroup
	topic    string
}

func ProviderListenerCustomerCreated(config *configuration.KafkaConfig) (*CustomerCreatedListener, error) {
	prop, err := properties.LoadFile("config.properties", properties.UTF8)

	if err != nil {
		panic(err)
	}

	return &CustomerCreatedListener{
		consumer: config.Consumer,
		topic:    prop.GetString("topic.request.process.evaluate", ""),
	}, nil
}

func (c *CustomerCreatedListener) Start() error {
	ctx := context.Background()

	for {
		if ctx.Err() != nil {
			return ctx.Err()
		}

		err := c.consumer.Consume(ctx, []string{c.topic}, c)
		if err != nil {
			return fmt.Errorf("erro ao consumir mensagem: %w", err)
		}
	}
}

func (c *CustomerCreatedListener) Setup(sarama.ConsumerGroupSession) error {
	log.Println("Setup listener")
	return nil
}

func (c *CustomerCreatedListener) Cleanup(sarama.ConsumerGroupSession) error {
	log.Println("finish consumer group...")
	return nil
}

func (c *CustomerCreatedListener) Close() error {
	if err := c.consumer.Close(); err != nil {
		return fmt.Errorf("erro close consumer group: %w", err)
	}
	log.Println("close consumer group...")
	return nil
}

func (c *CustomerCreatedListener) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		if err := c.handleMessage(message); err != nil {
			log.Printf("Erro ao processar mensagem: %v", err)
		}

		session.MarkMessage(message, "")
	}
	return nil
}

func (c *CustomerCreatedListener) handleMessage(message *sarama.ConsumerMessage) error {
	log.Printf("message received - Tópico: %s, Partição: %d, Offset: %d",
		message.Topic, message.Partition, message.Offset)

	var dto dtos.CustomerMessage
	if err := json.Unmarshal(message.Value, &dto); err != nil {
		return fmt.Errorf("error deserializer message: %w", err)
	}

	log.Printf("message process success - ID: %s", dto.Code)
	return nil
}

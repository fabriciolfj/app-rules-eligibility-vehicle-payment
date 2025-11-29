package adapters

import (
	"encoding/json"
	"log"

	"github.com/fabriciolfj/rules-elegibility/dtos"
	"github.com/fabriciolfj/rules-elegibility/entities"
	"github.com/fabriciolfj/rules-elegibility/producers"
	"github.com/magiconair/properties"
)

type NotifyCustomerCreatedAdapter struct {
	producer *producers.RuleProcessProducer
}

func ProviderNotifyCustomerCreatedAdapter(ruleProducer *producers.RuleProcessProducer) *NotifyCustomerCreatedAdapter {
	return &NotifyCustomerCreatedAdapter{
		producer: ruleProducer,
	}
}

func (notify NotifyCustomerCreatedAdapter) Process(customer *entities.Customer) {
	topic, _ := getTopic()
	dto := dtos.CustomerMessage{
		Code: customer.Code,
	}

	message, err := json.Marshal(dto)
	if err != nil {
		log.Fatal("error marshalling customer message %s", err)
	}

	notify.producer.SendMessage(string(message), topic)
}

func getTopic() (string, error) {
	prod, err := properties.LoadFile("config.properties", properties.UTF8)
	if err != nil {
		panic(err)
	}

	topic := prod.GetString("topic.request.process.evaluate", "")
	if topic == "" {
		log.Fatal("no topic information provided")
	}

	return topic, nil
}

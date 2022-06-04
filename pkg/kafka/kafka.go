package kafka

import "github.com/Shopify/sarama"

type Broker struct {
	Broker *sarama.Broker
	
}

func NewKafka(KafkaURL string) Broker {
	return Broker{
		Broker: sarama.NewBroker(KafkaURL),
	}
}

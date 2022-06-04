package helper

import (
	"encoding/json"

	"github.com/Shopify/sarama"
)

func Message(topic string, entity interface{}) (*sarama.ProducerMessage, error) {
	data, err := json.Marshal(entity)
	if err != nil {
		return nil, err
	}
	return &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(data),
	}, nil
}

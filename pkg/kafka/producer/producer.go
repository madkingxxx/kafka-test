package producer

import (
	"context"
	"kafka-pub-sub/pkg/entities"
	"kafka-pub-sub/pkg/helper"

	"github.com/Shopify/sarama"
	"go.uber.org/zap"
)

type Producer struct {
	producer sarama.SyncProducer
	*entities.ServiceOptions
}

func NewProducer(conf *sarama.Config, opt *entities.ServiceOptions) (*Producer, error) {
	producer, err := sarama.NewSyncProducer([]string{opt.Cfg.KafkaURL}, conf)
	if err != nil {
		return nil, err
	}
	return &Producer{
		producer:       producer,
		ServiceOptions: opt,
	}, nil
}

func (p *Producer) SendMessage(ctx context.Context, topic string, entity interface{}) error {
	p.Log.Info("Send message: ", zap.Any("message", entity))
	msg, err := helper.Message(topic, entity)
	if err != nil {
		return err
	}
	_, _, err = p.producer.SendMessage(msg)
	return err
}

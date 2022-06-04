package subscriber

import (
	"context"
	"log"

	"kafka-pub-sub/pkg/entities"

	"github.com/Shopify/sarama"
	"go.uber.org/zap"
)

type Consumer struct {
	consumer sarama.Consumer
	*entities.ServiceOptions
}

func NewSubscriber(conf *sarama.Config, opt *entities.ServiceOptions) *Consumer {
	consumer, err := sarama.NewConsumer([]string{opt.Cfg.KafkaURL}, conf)
	if err != nil {
		log.Fatal("error while NewSyncProducer: ", zap.Error(err))
		return nil
	}
	return &Consumer{
		consumer:       consumer,
		ServiceOptions: opt,
	}
}

func (s *Consumer) Listen(ctx context.Context, topic string) (err error) {
	s.Log.Info("listening to", zap.String("topic", topic))
	partitionConsumer, err := s.consumer.ConsumePartition(topic, 0, -1)
	if err != nil {
		return err
	}
	defer partitionConsumer.Close()
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			s.Log.Info("Received: ", zap.String("topic", topic), zap.String("value", string(msg.Value)))
		case err = <-partitionConsumer.Errors():
			s.Log.Error("Error: ", zap.Error(err))
			return err
		}
	}
	return nil
}

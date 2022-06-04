package main

import (
	"context"
	"kafka-pub-sub/config"
	"kafka-pub-sub/pkg/entities"
	"kafka-pub-sub/pkg/kafka/producer"
	"kafka-pub-sub/pkg/kafka/subscriber"
	"log"
	"time"

	"github.com/Shopify/sarama"
	"go.uber.org/zap"
)

func main() {
	cfg := config.NewConfig()
	logger, err := zap.NewProduction()
	if err != nil {
		log.Panic(err)
	}
	conf := sarama.NewConfig()
	conf.Producer.Return.Successes = true
	conf.Version = sarama.V2_0_0_0
	conf.Consumer.Offsets.Initial = sarama.OffsetOldest
	conf.Producer.Timeout = 5 * time.Second
	opt := &entities.ServiceOptions{
		Log: logger,
		Cfg: cfg,
	}
	producer, err := producer.NewProducer(conf, opt)
	if err != nil {
		log.Panic(err)
	}
	ctx := context.Background()
	topic := "bekzod"
	consumer := subscriber.NewSubscriber(conf, opt)
	go func() {
		err = consumer.Listen(ctx, topic)
		if err != nil {
			log.Panic(err)
		}
	}()
	err = producer.SendMessage(ctx, topic, "bekzod is nigger?!")
	if err != nil {
		log.Panic(err)
	}
	<-ctx.Done()
}

package kafkalog

import (
	"github.com/bsm/sarama-cluster"
	"github.com/Shopify/sarama"
	"log"
	"strings"
	config2 "gaecharge/config"
	"time"
)

func StartConsumer(msgFunc func(message *sarama.ConsumerMessage) error) error {
	config := cluster.NewConfig()
	config.Consumer.Offsets.CommitInterval = time.Second * 1
	config.Consumer.Offsets.Initial = sarama.OffsetNewest

	consumer, err := cluster.NewConsumer(
		strings.Split(
			config2.AppConfig.Kafka.BrokerList, ","),
			config2.AppConfig.Kafka.Group,
			[]string{config2.AppConfig.Kafka.Topic},
			config)

	if nil != err {
		return err
	}
	defer consumer.Close()

	log.Println("start listening")
	for msg := range consumer.Messages() {
		msgFunc(msg)
	}

	return nil
}

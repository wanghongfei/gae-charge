package kafkalog

import (
	"github.com/bsm/sarama-cluster"
	"github.com/Shopify/sarama"
	"log"
)

func StartConsumer(msgFunc func(message *sarama.ConsumerMessage) error) error {
	config := cluster.NewConfig()
	config.Consumer.Offsets.CommitInterval = 1
	config.Consumer.Offsets.Initial = sarama.OffsetNewest

	consumer, err := cluster.NewConsumer([]string{"10.150.182.11:8092"}, "gae-charge-1", []string{"gae-charge"}, config)
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

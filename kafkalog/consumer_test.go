package kafkalog

import (
	"testing"
	"github.com/Shopify/sarama"
	"fmt"
)

func TestStartConsumer(t *testing.T) {
	StartConsumer(func(msg *sarama.ConsumerMessage) error {
		fmt.Println(string(msg.Value))
		return nil
	})
}

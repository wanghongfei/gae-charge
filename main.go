package main

import (
	"gaecharge/kafkalog"
	"gaecharge/biz"
)

func main() {
	kafkalog.StartConsumer(biz.ConsumeMessage)
}

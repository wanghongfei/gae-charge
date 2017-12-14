package main

import (
	"gaecharge/kafkalog"
	"gaecharge/biz"
	"gaecharge/biz/report/cron"
	"gaecharge/biz/report"
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// kafak监听
	go func() {
		wg.Add(1)
		kafkalog.StartConsumer(biz.ConsumeMessage)
	}()

	// 报表定时任务
	go func() {
		wg.Add(1)

		for {
			ticker := cron.NewHourTicker(20, 1)
			ticker.Tick() // block

			err := report.CalculateHourlyReport()
			if nil != err {
				log.Println(err)
			}
		}
	}()

	wg.Wait()
}

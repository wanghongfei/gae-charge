package main

import (
	"gaecharge/kafkalog"
	"gaecharge/biz"
	"gaecharge/biz/report/cron"
	"gaecharge/biz/report"
	"log"
	"sync"
	"time"
	"gaecharge/config"
)

func main() {
	time.Sleep(time.Second * 10)

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
			ticker := cron.NewHourTicker(config.AppConfig.Report.ExeMinute, config.AppConfig.Report.ExeSecond)
			ticker.Tick() // block

			err := report.CalculateHourlyReport()
			if nil != err {
				log.Println(err)
			}
		}
	}()

	wg.Wait()
}

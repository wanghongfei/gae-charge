package main

import (
	"gaecharge/kafkalog"
	"gaecharge/biz"
	"gaecharge/biz/report/cron"
	"gaecharge/biz/report"
	"log"
	"sync"
	"gaecharge/config"
	_ "net/http/pprof"
	"net/http"
)

func main() {
	var wg sync.WaitGroup

	// kafak监听
	wg.Add(1)
	startKafkaConsumer()

	// 报表定时任务
	wg.Add(1)
	startReportTask()


	startProfiling()
	wg.Wait()
}

func startKafkaConsumer() {
	go func() {
		kafkalog.StartConsumer(biz.ConsumeMessage)
	}()
}

func startReportTask() {
	go func() {
		for {
			ticker := cron.NewHourTicker(config.AppConfig.Report.ExeMinute, config.AppConfig.Report.ExeSecond)
			ticker.Tick() // block

			err := report.CalculateHourlyReport()
			if nil != err {
				log.Println(err)
			}
		}
	}()

}

func startProfiling() {
	// pprof.StartCPUProfile(f)

	go func() {
		http.ListenAndServe("localhost:9000", nil)
	}()
}

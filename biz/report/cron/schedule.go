package cron

import (
	"time"
)

// 每小时定时tick
type HourTicker struct {
	minute	int
	second	int

	tick	chan interface{}
}

// 创建一个每小时定点ticker
// min: tick的分钟数
// sec: tick的秒数
func NewHourTicker(min, sec int) (*HourTicker) {
	if sec < 0 {
		sec = 0
	}

	if min < 0 {
		min = 0
	}

	return &HourTicker{
		minute: min,
		second: sec,
	}
}

func (obj *HourTicker) Tick() {
	// 计算下次执行时间
	waitSecond := 0
	now := time.Now()

	nowMin := now.Minute()
	nowSec := now.Second()
	nowTime := nowSec + nowMin * 60
	exeTime := obj.second + obj.minute * 60

	if nowTime < exeTime {
		waitSecond += exeTime - nowTime
	} else if nowTime == exeTime {
		waitSecond += 3600
	} else {
		waitSecond += 3600 - nowTime + exeTime
	}

	delay := int64(waitSecond * 1000)
	<- time.After(time.Millisecond * time.Duration(delay))
}


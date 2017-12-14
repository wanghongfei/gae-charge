package cron

import (
	"testing"
	"fmt"
	"time"
)

func TestScheduler_Tick(t *testing.T) {
	ticker, _ := NewHourTicker(24, 2)
	for {
		ticker.Tick()
		now := time.Now()

		fmt.Printf("%d:%d:%d\n", now.Hour(), now.Minute(), now.Second())
	}
}

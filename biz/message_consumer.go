package biz

import (
	"github.com/Shopify/sarama"
	"gaecharge/kafkalog"
	"gaecharge/biz/rediss"
	"time"
	"os"
	"log"
)

const KEY_CHARGE_PREFIX = "gae:charge:"
const SQL_OFFLINE_UNIT = "SELECT ?"

func ConsumeMessage(msg *sarama.ConsumerMessage) error {
	line := string(msg.Value)
	log.Println(line)

	chargeLog, err := kafkalog.ParseChargeLog(line)
	if nil != err {
		log.Println(err)
		return err
	}

	key := KEY_CHARGE_PREFIX + string(chargeLog.UnitId)
	cost := chargeLog.Bid / 1000

	// 扣费
	left, err := rediss.Charge(key, cost)
	if nil != err {
		log.Println(err)
		return err

	}

	// 落盘
	save(chargeLog.ExposeTime, line)


	// 下线
	if left <= 0 {

	}

	return nil
}

func save(exposeTime int64, data string) error {
	ts := time.Unix(exposeTime, 0)
	timeStr := ts.Format("2006010215")

	logFile, err := os.OpenFile("/tmp/" + timeStr + ".log", os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0644)
	if nil != err {
		return err
	}
	defer logFile.Close()

	_, err = logFile.Write([]byte(data + "\n"))
	if nil != err {
		return err
	}

	return nil
}

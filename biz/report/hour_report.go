package report

import (
	"time"
	"os"
	"bufio"
	"io"
	"strings"
	"gaecharge/kafkalog"
	"strconv"
	"gaecharge/config"
)

type ReportResult struct {
	UnitMap		map[string]int64
	IdeaMap		map[string]int64
	PlanMap		map[string]int64

	SlotMap		map[string]int64
	RegionMap	map[string]int64
	TagMap		map[string]int64
}

func CalculateHourlyReport() error {
	// 上个小时的时间日志文件
	filename := config.AppConfig.Report.InputDir + "/" + getLastHour() + ".log"

	result := &ReportResult{
		UnitMap: make(map[string]int64),
		IdeaMap: make(map[string]int64),
		PlanMap: make(map[string]int64),

		SlotMap: make(map[string]int64),
		RegionMap: make(map[string]int64),
		TagMap: make(map[string]int64),
	}
	readFileLineByLine(filename, result, processLine)

	writeFile(result)

	return nil
}

func writeFile(result *ReportResult) error {
	writeMap(".unit", result.UnitMap)
	writeMap(".idea", result.IdeaMap)
	writeMap(".plan", result.PlanMap)
	writeMap(".slot", result.SlotMap)
	writeMap(".region", result.RegionMap)
	writeMap(".tag", result.TagMap)

	return nil
}

func writeMap(suffix string, dataMap map[string]int64) error {
	prefix := config.AppConfig.Report.OutputDir + "/" + getLastHour() + ".log"
	f, err := os.Create(prefix + suffix)
	if nil != err {
		return err
	}
	defer f.Close()

	for k, v := range dataMap {
		data := k + "\t" + strconv.FormatInt(v, 10) + "\n"
		f.WriteString(data)
	}

	return nil
}

func processLine(line string, result *ReportResult) {
	chargeLog, err := kafkalog.ParseChargeLog(line)
	if nil != err {
		return
	}

	result.PlanMap[strconv.Itoa(chargeLog.PlanId)]++
	result.UnitMap[strconv.Itoa(chargeLog.UnitId)]++
	result.IdeaMap[chargeLog.IdeaId]++

	result.SlotMap[chargeLog.SlotCode]++
	result.RegionMap[strconv.Itoa(chargeLog.RegionId)]++
	result.TagMap[chargeLog.TagIds]++
}

func getLastHour() string {
	now := time.Now()
	diff, _:= time.ParseDuration("-1h")
	hourAgo := now.Add(diff)
	timeString := hourAgo.Format("2006010215")

	return timeString
}

func readFileLineByLine(filename string, result *ReportResult, lineFunc func(string, *ReportResult)) error {
	logFile, err := os.Open(filename)
	if nil != err {
		return err
	}
	defer logFile.Close()

	reader := bufio.NewReader(logFile)
	for {
		line, err := reader.ReadString('\n')
		if nil != err {
			if io.EOF == err {
				break
			}

			return err
		}

		lineFunc(strings.TrimSpace(line), result)
	}

	return nil
}
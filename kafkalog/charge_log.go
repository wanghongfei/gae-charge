package kafkalog

import (
	"fmt"
	"strings"
	"strconv"
)

type ChargeLog struct {
	ExposeTime		int64
	Sid				string
	SearchTime		int64

	Tid				string
	RequestId		string
	ResourceType	int
	SlotCode		string
	SlotType		int
	Width			int
	Height			int
	MaterialType	int
	Mac				string
	Ip				string

	PlanId			int
	UnitId			int
	IdeaId			string
	AdCode			string
	TagIds			string
	RegionId		int
	Bid				int64
}

// 将日志字符串解析成日志对象
func ParseChargeLog(str string) (*ChargeLog, error) {
	terms := strings.Split(str, "\t")
	if nil == terms || len(terms) != 21 {
		return nil, fmt.Errorf("invalid line: %s", str)
	}

	exposeTime, err := strconv.ParseInt(terms[1], 10, 64)
	if nil != err {
		return nil, err
	}

	sid := terms[2]

	searchTime, err := strconv.ParseInt(terms[3], 10, 64)
	if nil != err {
		return nil, err
	}

	tid := terms[4]
	requestId := terms[5]

	resourceType, err := strconv.Atoi(terms[6])
	if nil != err {
		return nil, err
	}

	slotCode := terms[7]

	slotType, err := strconv.Atoi(terms[8])
	if nil != err {
		return nil, err
	}

	width, err := strconv.Atoi(terms[9])
	if nil != err {
		return nil, err
	}

	height, err := strconv.Atoi(terms[10])
	if nil != err {
		return nil, err
	}

	materialType, err := strconv.Atoi(terms[11])
	if nil != err {
		return nil, err
	}

	mac := terms[12]
	ip := terms[13]


	planId, err := strconv.Atoi(terms[14])
	if nil != err {
		return nil, err
	}

	unitId, err := strconv.Atoi(terms[15])
	if nil != err {
		return nil, err
	}

	ideaId := terms[16]
	adCode := terms[17]
	tagIds := terms[18]

	regionId, err := strconv.Atoi(terms[19])
	if nil != err {
		return nil, err
	}

	bid, err := strconv.ParseInt(terms[20], 10, 64)
	if nil != err {
		return nil, err
	}

	return &ChargeLog{
		ExposeTime: exposeTime,
		Sid: sid,
		SearchTime: searchTime,
		Tid: tid,
		RequestId: requestId,
		ResourceType: resourceType,
		SlotCode: slotCode,
		SlotType: slotType,
		Width: width,
		Height: height,
		MaterialType: materialType,
		Mac: mac,
		Ip: ip,
		PlanId: planId,
		UnitId: unitId,
		IdeaId: ideaId,
		AdCode: adCode,
		TagIds: tagIds,
		RegionId: regionId,
		Bid: bid,
	}, nil
}

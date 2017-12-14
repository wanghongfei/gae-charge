package biz

import (
	"testing"
	"time"
)

func TestSave(t *testing.T) {
	err := save(time.Now().Unix(), "data")
	if nil != err {
		t.Errorf("%v\n", err)
	}
}

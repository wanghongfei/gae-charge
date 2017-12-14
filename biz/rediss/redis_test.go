package rediss

import (
	"testing"
	"fmt"
)

func TestCharge(t *testing.T) {
	left, err := Charge("gae:charge:100", 5)
	if nil != err {
		t.Errorf("%v\n", err)
		return
	}

	fmt.Println(left)
}

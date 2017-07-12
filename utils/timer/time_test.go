package timer

import (
	"testing"
	"time"
)

func Test_timeAlertOk(t *testing.T) {

	b := AlertTime("00:00", "23:59", "0,1,2,3,4,5,6", time.Now())

	if b == false {
		t.Fatal("error alert time")
	}
}

func Test_timeAlertFail(t *testing.T) {
	b := AlertTime("00:00", "00:01", "0,1,2,3,4,5,6", time.Now())

	if b == true {
		t.Fatal("error alert time")
	}
}

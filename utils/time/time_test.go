package time

import (
	"testing"
	"time"
)

func Test_timeAlertOk(t *testing.T) {

	var tc []int64 = []int64{0, 1, 2, 3, 4, 5, 6}

	b := AlertTime("00:00", "23:59", tc, time.Now())

	if b == false {
		t.Fatal("error alert time")
	}
}

func Test_timeAlertFail(t *testing.T) {

	var tc []int64 = []int64{0, 1, 2, 3, 4, 5, 6}

	b := AlertTime("00:00", "00:01", tc, time.Now())

	if b == true {
		t.Fatal("error alert time")
	}
}

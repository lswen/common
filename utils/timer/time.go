package timer

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

var DAYS = map[string]int64{
	"Sunday":    0,
	"Monday":    1,
	"Tuesday":   2,
	"Wednesday": 3,
	"Thursday":  4,
	"Friday":    5,
	"Saturday":  6,
}

func AlertTime(start string, end string, strday string, now time.Time) bool {

	day := stringtoArrayInt(strday)

	if start == "" || end == "" || len(day) == 0 {
		return false
	}

	startTimeStr := strings.Split(start, ":")

	if len(startTimeStr) != 2 {
		return false
	}

	endTimeStr := strings.Split(end, ":")

	if len(endTimeStr) != 2 {
		return false
	}

	startHour, err := strconv.Atoi(startTimeStr[0])

	if err != nil {
		return false
	}

	startMinute, err := strconv.Atoi(startTimeStr[1])

	if err != nil {
		return false
	}

	endMinute, err := strconv.Atoi(endTimeStr[1])

	if err != nil {
		return false
	}

	endHour, err := strconv.Atoi(endTimeStr[0])

	if err != nil {
		return false
	}

	//currHour := now.Hour()
	//currMinute := now.Minute()
	currDay := now.Weekday().String()

	math := false
	for i := 0; i < len(day); i++ {

		if day[i] == DAYS[currDay] {
			math = true
			break
		}
	}

	if math == false {
		log.Println("星期不通过")
		return false
	}

	sts := fmt.Sprintf("%04d-%02d-%02d %02d:%02d:00", now.Year(), now.Month(), now.Day(), startHour, startMinute)
	ets := fmt.Sprintf("%04d-%02d-%02d %02d:%02d:00", now.Year(), now.Month(), now.Day(), endHour, endMinute)

	su := TimeStringToTime(sts).Unix()
	eu := TimeStringToTime(ets).Unix()

	if su <= now.Unix() && eu >= now.Unix() {
		return true
	}

	log.Println("no match", su, eu, now.Unix(), sts, ets)

	return false
}
func stringtoArrayInt(str string) (arrayInt []int64) {

	arrStr := strings.Split(str, ",")

	for i := 0; i < len(arrStr); i++ {

		data, err := strconv.ParseInt(arrStr[i], 10, 64)

		if err != nil {
			log.Println(data, err)
			continue
		}

		arrayInt = append(arrayInt, data)
	}

	return
}

func TimeStringToTime(t string) time.Time {
	//	toBeCharge := "2015-01-01 00:00:00"
	layout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	ct, _ := time.ParseInLocation(layout, t, loc)
	return ct
}

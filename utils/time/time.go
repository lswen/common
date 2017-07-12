package time

import (
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

func AlertTime(start string, end string, day []int64, now time.Time) bool {

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

	currHour := now.Hour()
	currMinute := now.Minute()
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

	if startHour <= currHour && endHour >= currHour {

		if startHour == currHour {
			//判断分钟
			if startMinute <= currMinute {
				log.Println("同开始时间时间到:", start, end, day)
				//continue
			} else {
				log.Println("1不在时间段:", start, end, day)
				return false
			}
		} else if endHour == currHour {
			if endMinute >= currMinute {
				log.Println("同结束时间，时间到:", start, end, day)
				//continue
			} else {
				log.Println("2不在时间段:", start, end, day)
				return false
			}
		} else {
			log.Println("else 时间到:", start, end, day)
			//continue
		}

	} else {
		log.Println("3不在时间段:", start, end, day, startHour, currHour, endHour, currHour)
		return false
	}

	return true
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

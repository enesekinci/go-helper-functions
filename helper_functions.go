package helper_functions

import (
	"strconv"
	"strings"
	"time"
)

func dateConvert(date time.Time, format string) string {

	day := strconv.Itoa(date.Day())
	if len(day) == 1 {
		day = "0" + day
	}

	month := strconv.Itoa(int(date.Month()))

	if month != "0" && len(month) == 1 {
		month = "0" + month
	} else if month == "0" {
		panic("invalid month")
	}
	year := strconv.Itoa(date.Year())

	hour := strconv.Itoa(date.Hour())
	if len(hour) == 1 {
		hour = "0" + hour
	}
	minute := strconv.Itoa(date.Minute())
	if len(minute) == 1 {
		minute = "0" + minute
	}
	second := strconv.Itoa(date.Second())
	if len(second) == 1 {
		second = "0" + second
	}

	format = strings.Replace(format, "d", day, 1)
	format = strings.Replace(format, "m", month, 1)
	format = strings.Replace(format, "Y", year, 1)
	format = strings.Replace(format, "H", hour, 1)
	format = strings.Replace(format, "i", minute, 1)
	format = strings.Replace(format, "s", second, 1)

	return format
}

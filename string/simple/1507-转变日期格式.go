package simple

import (
	"strings"
)

func reformatDate(date string) string {
	monthMap := map[string]string{
		"Jan": "01", "Feb": "02", "Mar": "03", "Apr": "04", "May": "05", "Jun": "06",
			"Jul": "07", "Aug": "08", "Sep": "09", "Oct": "10", "Nov": "11", "Dec": "12",
	}
	dates := strings.Split(date,  " ")

	return dates[2] + "-" + monthMap[dates[1]] + "-" + handleDay(dates[0])
}

func handleDay(day  string) string {
	sub :=  day[0:len(day)-2]
	if len(sub) == 1 {
		sub  = "0" + sub
	}

	return sub
}


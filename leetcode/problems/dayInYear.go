package problems

import (
	"fmt"
	"time"
)

//输入某年某月某日，判断这一天是这一年的第几天？
func dayInYear(year, month, day int) (result int) {
	if month < 1 || month > 12 {
		panic("month must more than 0 and less than 13!")
	}

	if day < 1 || day > 31 {
		panic("day must more than 0 and less than 32!")
	}

	if month == 1 {
		return day
	}
	monthDay := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		monthDay[1] = 29
	}

	monthMaxDay := monthDay[month-1]
	if day > monthMaxDay {
		message := fmt.Sprintf("%v of year %d has %d days,you cannot input a number more %d as day number!", time.Month(month), year, monthMaxDay, monthMaxDay)
		panic(message)
	}

	for i := 1; i < month; i++ {
		result += monthDay[i-1]
	}

	return result + day
}

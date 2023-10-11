package day1

import (
	"fmt"
	"strconv"
)

// https://www.hackerrank.com/challenges/one-week-preparation-kit-time-conversion

func timeConversion(s string) {
	hour, _ := strconv.Atoi(s[0:2])
	if s[8:9] == "P" && hour < 12 {
		hour += 12
	} else if s[8:9] == "A" && hour == 12 {
		hour = 0
	}
	fmt.Printf("%02d%s", hour, s[2:8])
}

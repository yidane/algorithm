package others

import (
	"strconv"
)

/*
Given an integer, return its base 7 string representation.

Example 1:
Input: 100
Output: "202" 7^2*2+7^1*0+7^0*2=100
Example 2:
Input: -7
Output: "-10"
Note: The input will be in range of [-1e7, 1e7].
*/

func convertToBase7(num int) string {
	if num > -7 && num < 7 {
		return strconv.Itoa(num)
	}

	maxBit := 7
	maxStep := 1
	b := false
	if num < 0 {
		b = true
		num = -num
	}
	for num/maxBit >= 7 {
		maxBit = maxBit * 7
		maxStep++
	}

	outString := ""
	if b {
		outString += "-"
	}
	for i := maxStep; i >= 0; i-- {
		t := 0
		if num >= maxBit {
			t = num / maxBit
		}
		outString += strconv.Itoa(t)
		num = num - maxBit*t
		maxBit = maxBit / 7
	}

	return outString
}

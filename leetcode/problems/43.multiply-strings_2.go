package problems

import (
	"bytes"
)

/*
 * @lc app=leetcode id=43 lang=golang
 *
 * [43] Multiply Strings
 *
 * https://leetcode.com/problems/multiply-strings/description/
 *
 * algorithms
 * Medium (30.04%)
 * Total Accepted:    186K
 * Total Submissions: 619.2K
 * Testcase Example:  '"2"\n"3"'
 *
 * Given two non-negative integers num1 and num2 represented as strings, return
 * the product of num1 and num2, also represented as a string.
 *
 * Example 1:
 *
 *
 * Input: num1 = "2", num2 = "3"
 * Output: "6"
 *
 * Example 2:
 *
 *
 * Input: num1 = "123", num2 = "456"
 * Output: "56088"
 *
 *
 * Note:
 *
 *
 * The length of both num1 and num2 is < 110.
 * Both num1 and num2 contain only digits 0-9.
 * Both num1 and num2 do not contain any leading zero, except the number 0
 * itself.
 * You must not use any built-in BigInteger library or convert the inputs to
 * integer directly.
 *
 *
 */
/*
		  1  2  3
		* 4  5  6
		  6  12	18
	   5  10 15
  + 4  8  12
	5  6  0  8  8

len(m)+len(n)-1  [len(m)][len(m)+len(n)-1]arr
*/

//todo not the best
func multiply(num1 string, num2 string) string {
	width := len(num1) + len(num2) - 1
	arr := make(map[int]map[int]int, width)

	for i := 0; i < len(num1); i++ {
		subArr := make(map[int]int)
		for j := 0; j < len(num2); j++ {
			subArr[i+j] = int(num1[i]-48) * int(num2[j]-48)
		}
		arr[i] = subArr
	}

	addBit := 0

	arr1 := make(map[int]int, width)
	for i := width - 1; i >= 0; i-- {
		t := 0
		for j := 0; j < len(arr); j++ {
			t += arr[j][i]
		}
		t += addBit

		addBit = t / 10
		arr1[i] = t % 10
	}

	buf := bytes.Buffer{}
	if addBit > 0 {
		buf.WriteRune(rune(addBit + 48))
	}

	for i := 0; i < len(arr1); i++ {
		r := arr1[i]
		if buf.Len() == 0 && r == 0 {
			continue
		}
		buf.WriteRune(rune(r + 48))
	}

	if buf.Len() == 0 {
		return "0"
	}

	return buf.String()
}

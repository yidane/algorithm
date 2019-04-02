package problems

import (
	"bytes"
)

/*
 * @lc app=leetcode id=6 lang=golang
 *
 * [6] ZigZag Conversion
 *
 * https://leetcode.com/problems/zigzag-conversion/description/
 *
 * algorithms
 * Medium (31.07%)
 * Total Accepted:    299.5K
 * Total Submissions: 962.3K
 * Testcase Example:  '"PAYPALISHIRING"\n3'
 *
 * The string "PAYPALISHIRING" is written in a zigzag pattern on a given number
 * of rows like this: (you may want to display this pattern in a fixed font for
 * better legibility)
 *
 *
 * P   A   H   N
 * A P L S I I G
 * Y   I   R
 *
 *
 * And then read line by line: "PAHNAPLSIIGYIR"
 *
 * Write the code that will take a string and make this conversion given a
 * number of rows:
 *
 *
 * string convert(string s, int numRows);
 *
 * Example 1:
 *
 *
 * Input: s = "PAYPALISHIRING", numRows = 3
 * Output: "PAHNAPLSIIGYIR"
 *
 *
 * Example 2:
 *
 *
 * Input: s = "PAYPALISHIRING", numRows = 4
 * Output: "PINALSIGYAHRPI"
 * Explanation:
 *
 * P     I    N
 * A   L S  I G
 * Y A   H R
 * P     I
 *
 */

/*
肯定可以用数学公式表达出来
PAYPALISHIRING
PINALSIGYAHRPI

123456789
172683594

*/

func convert(s string, numRows int) string {
	if numRows < 2 || len(s) < numRows {
		return s
	}

	arr := make([]*bytes.Buffer, 0)

	for i := 0; i < numRows; i++ {
		arr = append(arr, bytes.NewBufferString(s[i:i+1]))
	}

	var f bool
	for i, j := numRows, numRows-2; i < len(s); i++ { //定义两个游标，一个对字符串递增处理，一个控制折线方向
		arr[j].WriteRune(rune(s[i]))

		if j == 0 || j == numRows-1 {
			f = !f
		}

		if f {
			j++
		} else {
			j--
		}
	}

	buf := bytes.Buffer{}

	for i := 0; i < len(arr); i++ {
		buf.Write(arr[i].Bytes())
	}

	return buf.String()
}

/*
	if len(s) < numRows {
		return s
	}

	result := make([]uint8, len(s))

	partSize := numRows + numRows - 2
	totalPart := len(s) / partSize
	for i := 0; i < totalPart*partSize; i++ {
		m := i / partSize
		newIndex := 0
		switch i % partSize {
		case 0:
			newIndex = m
			result[newIndex] = s[i]
		case numRows - 1:

		default:

		}
	}

	return string(result)
*/

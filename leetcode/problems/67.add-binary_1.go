package problems

import (
	"bytes"
)

/*
 * @lc app=leetcode id=67 lang=golang
 *
 * [67] Add Binary
 *
 * https://leetcode.com/problems/add-binary/description/
 *
 * algorithms
 * Easy (38.62%)
 * Total Accepted:    292.7K
 * Total Submissions: 757K
 * Testcase Example:  '"11"\n"1"'
 *
 * Given two binary strings, return their sum (also a binary string).
 *
 * The input strings are both non-empty and contains only characters 1 or 0.
 *
 * Example 1:
 *
 *
 * Input: a = "11", b = "1"
 * Output: "100"
 *
 * Example 2:
 *
 *
 * Input: a = "1010", b = "1011"
 * Output: "10101"
 *
 */
func addBinary(a string, b string) string {
	buffer := bytes.Buffer{}
	addBit := false

	for i, j := len(a)-1, len(b)-1; i > -1 || j > -1; {
		var r uint8 = '0'
		if i == -1 {
			r, addBit = addBinaryFunc(r, b[j], addBit)
			buffer.WriteByte(r)
			j--
			continue
		}
		if j == -1 {
			r, addBit = addBinaryFunc(a[i], r, addBit)
			buffer.WriteByte(r)
			i--
			continue
		}
		r, addBit = addBinaryFunc(a[i], b[j], addBit)
		buffer.WriteByte(r)
		i--
		j--
	}

	if addBit {
		buffer.WriteRune('1')
	}

	bufBytes := buffer.Bytes()
	l := len(bufBytes)
	result := make([]byte, l)

	for i := l - 1; i >= 0; i-- {
		result[l-1-i] = bufBytes[i]
	}

	return string(result)
}

func addBinaryFunc(a, b uint8, addBit bool) (uint8, bool) {
	if addBit {
		if a == '0' {
			if b == '1' {
				return '0', true
			}
			return '1', false
		}

		if b == '0' {
			if a == '1' {
				return '0', true
			}
			return '1', false
		}

		return '1', true
	}

	if a == '0' {
		return b, false
	}

	if b == '0' {
		return a, false
	}

	return '0', true
}

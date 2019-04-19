package problems

import (
	"bytes"
	"strconv"
)

/*
 * @lc app=leetcode id=166 lang=golang
 *
 * [166] Fraction to Recurring Decimal
 *
 * https://leetcode.com/problems/fraction-to-recurring-decimal/description/
 *
 * algorithms
 * Medium (19.33%)
 * Total Accepted:    85.7K
 * Total Submissions: 443K
 * Testcase Example:  '1\n2'
 *
 * Given two integers representing the numerator and denominator of a fraction,
 * return the fraction in string format.
 *
 * If the fractional part is repeating, enclose the repeating part in
 * parentheses.
 *
 * Example 1:
 *
 *
 * Input: numerator = 1, denominator = 2
 * Output: "0.5"
 *
 *
 * Example 2:
 *
 *
 * Input: numerator = 2, denominator = 1
 * Output: "2"
 *
 * Example 3:
 *
 *
 * Input: numerator = 2, denominator = 3
 * Output: "0.(6)"
 *
 *
 */
func fractionToDecimal(numerator int, denominator int) string {
	if numerator%denominator == 0 {
		return strconv.Itoa(numerator / denominator)
	}

	f := (numerator < 0 && denominator > 0) || (numerator > 0 && denominator < 0)

	numLeft := numerator / denominator
	numOther := numerator % denominator
	result := bytes.Buffer{}
	if f {
		numOther = -numOther
		result.WriteString("-")
		numLeft = -numLeft
	}

	result.WriteString(strconv.Itoa(numLeft))

	result.WriteString(".")
	m := make(map[int]int)
	arr := make([]int, 0)

	i := -1
	ok := false
	for numOther != 0 {
		i, ok = m[numOther]
		if ok {
			break
		} else {
			i = -1
			m[numOther] = len(arr)
		}
		numOther = numOther * 10
		arr = append(arr, numOther/denominator)
		numOther = numOther % denominator
	}

	for k := 0; k < len(arr); k++ {
		if k == i {
			result.WriteString("(")
		}
		result.WriteString(strconv.Itoa(arr[k]))
	}
	if i > -1 {
		result.WriteString(")")
	}

	return result.String()
}

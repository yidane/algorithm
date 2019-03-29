package problems

import (
	"bytes"
)

/*
 * @lc app=leetcode id=168 lang=golang
 *
 * [168] Excel Sheet Column Title
 *
 * https://leetcode.com/problems/excel-sheet-column-title/description/
 *
 * algorithms
 * Easy (28.65%)
 * Total Accepted:    166.2K
 * Total Submissions: 580.3K
 * Testcase Example:  '1'
 *
 * Given a positive integer, return its corresponding column title as appear in
 * an Excel sheet.
 *
 * For example:
 *
 *
 * ⁠   1 -> A
 * ⁠   2 -> B
 * ⁠   3 -> C
 * ⁠   ...
 * ⁠   26 -> Z
 * ⁠   27 -> AA
 * ⁠   28 -> AB
 * ⁠   ...
 *
 *
 * Example 1:
 *
 *
 * Input: 1
 * Output: "A"
 *
 *
 * Example 2:
 *
 *
 * Input: 28
 * Output: "AB"
 *
 *
 * Example 3:
 *
 *
 * Input: 701
 * Output: "ZY"
 *
 */
func convertToTitle(n int) string {
	var rs []rune

	for n > 26 {
		d := n % 26
		if d == 0 { //要考虑到刚好是26的整数倍，52=AZ...
			rs = append(rs, 'Z')
			n = (n - 1) / 26
		} else {
			rs = append(rs, rune(d+64))
			n = n / 26
		}
	}
	rs = append(rs, rune(n+64))

	buf := bytes.Buffer{}
	for i := len(rs) - 1; i >= 0; i-- {
		buf.WriteRune(rs[i])
	}

	return buf.String()
}

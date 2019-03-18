package problems

/*
 * @lc app=leetcode id=921 lang=golang
 *
 * [921] Minimum Add to Make Parentheses Valid
 *
 * https://leetcode.com/problems/minimum-add-to-make-parentheses-valid/description/
 *
 * algorithms
 * Medium (69.52%)
 * Total Accepted:    19.3K
 * Total Submissions: 27.7K
 * Testcase Example:  '"())"'
 *
 * Given a string S of '(' and ')' parentheses, we add the minimum number of
 * parentheses ( '(' or ')', and in any positions ) so that the resulting
 * parentheses string is valid.
 *
 * Formally, a parentheses string is valid if and only if:
 *
 *
 * It is the empty string, or
 * It can be written as AB (A concatenated with B), where A and B are valid
 * strings, or
 * It can be written as (A), where A is a valid string.
 *
 *
 * Given a parentheses string, return the minimum number of parentheses we must
 * add to make the resulting string valid.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: "())"
 * Output: 1
 *
 *
 *
 * Example 2:
 *
 *
 * Input: "((("
 * Output: 3
 *
 *
 *
 * Example 3:
 *
 *
 * Input: "()"
 * Output: 0
 *
 *
 *
 * Example 4:
 *
 *
 * Input: "()))(("
 * Output: 4
 *
 *
 *
 *
 *
 *
 * Note:
 *
 *
 * S.length <= 1000
 * S only consists of '(' and ')' characters.
 *
 *
 *
 *
 *
 *
 *
 *
 *
 */
func minAddToMakeValid(S string) int {
	if len(S) == 0 {
		return 0
	}

	var leftCount, rightCount int

	for _, r := range S {
		if r == '(' {
			leftCount++
		} else {
			if leftCount > 0 {
				leftCount-- //能凑一对，就少一对
			} else {
				rightCount++
			}
		}
	}

	return rightCount + leftCount
}

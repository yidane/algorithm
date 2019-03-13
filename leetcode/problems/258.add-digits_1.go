package problems

/*
 * @lc app=leetcode id=258 lang=golang
 *
 * [258] Add Digits
 *
 * https://leetcode.com/problems/add-digits/description/
 *
 * algorithms
 * Easy (53.58%)
 * Total Accepted:    228.6K
 * Total Submissions: 426.7K
 * Testcase Example:  '38'
 *
 * Given a non-negative integer num, repeatedly add all its digits until the
 * result has only one digit.
 *
 * Example:
 *
 *
 * Input: 38
 * Output: 2
 * Explanation: The process is like: 3 + 8 = 11, 1 + 1 = 2.
 * Since 2 has only one digit, return it.
 *
 *
 * Follow up:
 * Could you do it without any loop/recursion in O(1) runtime?
 */
func addDigits(num int) int {
	for num > 9 {
		n := 0
		for num > 0 {
			n += num % 10
			num /= 10
		}
		num = n
	}

	return num
}

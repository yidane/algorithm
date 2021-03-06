package problems

/*
 * @lc app=leetcode id=172 lang=golang
 *
 * [172] Factorial Trailing Zeroes
 *
 * https://leetcode.com/problems/factorial-trailing-zeroes/description/
 *
 * algorithms
 * Easy (37.31%)
 * Total Accepted:    152.9K
 * Total Submissions: 409.7K
 * Testcase Example:  '3'
 *
 * Given an integer n, return the number of trailing zeroes in n!.
 *
 * Example 1:
 *
 *
 * Input: 3
 * Output: 0
 * Explanation: 3! = 6, no trailing zero.
 *
 * Example 2:
 *
 *
 * Input: 5
 * Output: 1
 * Explanation: 5! = 120, one trailing zero.
 *
 * Note: Your solution should be in logarithmic time complexity.
 *
 */

//主要计算出现多少个5和10，100,1000...
func trailingZeroes(n int) int {
	count := 0
	for i := 5; i <= n; i += 5 {
		count++
		j := i / 5

		for j%5 == 0 {
			j = j / 5
			count++
		}
	}

	return count
}

//Best
//应该是规律计算出来的
func trailingZeroes1(n int) int {
	res := 0
	for n >= 5 {
		n /= 5
		res += n
	}
	return res
}

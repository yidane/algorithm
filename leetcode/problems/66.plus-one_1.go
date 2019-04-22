package problems

/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 *
 * https://leetcode.com/problems/plus-one/description/
 *
 * algorithms
 * Easy (40.99%)
 * Total Accepted:    377.1K
 * Total Submissions: 919.5K
 * Testcase Example:  '[1,2,3]'
 *
 * Given a non-empty array of digits representing a non-negative integer, plus
 * one to the integer.
 *
 * The digits are stored such that the most significant digit is at the head of
 * the list, and each element in the array contain a single digit.
 *
 * You may assume the integer does not contain any leading zero, except the
 * number 0 itself.
 *
 * Example 1:
 *
 *
 * Input: [1,2,3]
 * Output: [1,2,4]
 * Explanation: The array represents the integer 123.
 *
 *
 * Example 2:
 *
 *
 * Input: [4,3,2,1]
 * Output: [4,3,2,2]
 * Explanation: The array represents the integer 4321.
 *
 */
func plusOne(digits []int) []int {
	b := 0

	last := digits[len(digits)-1] + 1
	if last == 10 {
		last = 0
		b = 1
	}
	digits[len(digits)-1] = last

	for i := len(digits) - 2; i >= 0; i-- {
		if b == 0 {
			break
		}
		last = digits[i] + 1
		if last == 10 {
			digits[i] = 0
			continue
		}

		digits[i] = last
		b = 0
		break
	}

	if b == 0 {
		return digits
	}

	result := []int{b}
	result = append(result, digits...)
	return result
}

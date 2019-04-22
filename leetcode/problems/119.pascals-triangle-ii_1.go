package problems

/*
 * @lc app=leetcode id=119 lang=golang
 *
 * [119] Pascal's Triangle II
 *
 * https://leetcode.com/problems/pascals-triangle-ii/description/
 *
 * algorithms
 * Easy (42.90%)
 * Total Accepted:    196.2K
 * Total Submissions: 456.6K
 * Testcase Example:  '3'
 *
 * Given a non-negative index k where k ≤ 33, return the k^th index row of the
 * Pascal's triangle.
 *
 * Note that the row index starts from 0.
 *
 *
 * In Pascal's triangle, each number is the sum of the two numbers directly
 * above it.
 *
 * Example:
 *
 *
 * Input: 3
 * Output: [1,3,3,1]
 *
 *
 * Follow up:
 *
 * Could you optimize your algorithm to use only O(k) extra space?
 *
 */
func getRow(rowIndex int) []int {
	if rowIndex < 0 {
		return nil
	}

	if rowIndex == 0 {
		return []int{1}
	}

	if rowIndex == 1 {
		return []int{1, 1}
	}

	result := []int{1, 1}
	for i := 2; i <= rowIndex; i++ {
		newArr := make([]int, len(result)+1)
		newArr[0] = 1

		for j := 0; j < len(result)-1; j++ {
			newArr[j+1] = result[j] + result[j+1]
		}
		newArr[len(newArr)-1] = 1
		result = newArr
	}

	return result
}

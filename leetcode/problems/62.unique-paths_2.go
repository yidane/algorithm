package problems

/*
 * @lc app=leetcode id=62 lang=golang
 *
 * [62] Unique Paths
 *
 * https://leetcode.com/problems/unique-paths/description/
 *
 * algorithms
 * Medium (46.81%)
 * Total Accepted:    270.4K
 * Total Submissions: 577.3K
 * Testcase Example:  '3\n2'
 *
 * A robot is located at the top-left corner of a m x n grid (marked 'Start' in
 * the diagram below).
 *
 * The robot can only move either down or right at any point in time. The robot
 * is trying to reach the bottom-right corner of the grid (marked 'Finish' in
 * the diagram below).
 *
 * How many possible unique paths are there?
 *
 *
 * Above is a 7 x 3 grid. How many possible unique paths are there?
 *
 * Note: m and n will be at most 100.
 *
 * Example 1:
 *
 *
 * Input: m = 3, n = 2
 * Output: 3
 * Explanation:
 * From the top-left corner, there are a total of 3 ways to reach the
 * bottom-right corner:
 * 1. Right -> Right -> Down
 * 2. Right -> Down -> Right
 * 3. Down -> Right -> Right
 *
 *
 * Example 2:
 *
 *
 * Input: m = 7, n = 3
 * Output: 28
 *
 */

/*
	总共会行走m+n-2步，其中向右行走m-1步，向下走n-1步
	那么就是从m+n-2中选取m-1的排列组合
	  m-1
	(
	 m+n-2
*/
func uniquePaths(m int, n int) int {
	t1, t2 := 1, 1

	t := m
	if m < n {
		t = n
	}
	for i := t; i <= m+n-2; i++ {
		t1 *= i
	}

	for i := 1; i <= m+n-1-t; i++ {
		t2 *= i
	}

	return t1 / t2
}

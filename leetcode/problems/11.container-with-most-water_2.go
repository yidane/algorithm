package problems

/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 *
 * https://leetcode.com/problems/container-with-most-water/description/
 *
 * algorithms
 * Medium (42.12%)
 * Total Accepted:    328K
 * Total Submissions: 763.8K
 * Testcase Example:  '[1,8,6,2,5,4,8,3,7]'
 *
 * Given n non-negative integers a1, a2, ..., an , where each represents a
 * point at coordinate (i, ai). n vertical lines are drawn such that the two
 * endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together
 * with x-axis forms a container, such that the container contains the most
 * water.
 *
 * Note: You may not slant the container and n is at least 2.
 *
 *
 *
 *
 *
 * The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In
 * this case, the max area of water (blue section) the container can contain is
 * 49.
 *
 *
 *
 * Example:
 *
 *
 * Input: [1,8,6,2,5,4,8,3,7]
 * Output: 49
 *
 */

/*
 0 0 0
 1 1 1
 1 2 3 ; 1 6 6
 1 3 3 ; 2 2 2 ; 1 2 2
 1 5 5 ; 4 4 16 ; 5 2 10 ; 2 1 2
 1 4
*/
func maxArea(height []int) int {
	if len(height) <= 1 {
		return 0
	}

	max := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			x := j - i
			y := 0
			if height[i] > height[j] {
				y = height[j]
			} else {
				y = height[i]
			}

			m := x * y
			if max < m {
				max = m
			}
		}
	}

	return max
}

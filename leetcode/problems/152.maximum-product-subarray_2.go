package problems

/*
 * @lc app=leetcode id=152 lang=golang
 *
 * [152] Maximum Product Subarray
 *
 * https://leetcode.com/problems/maximum-product-subarray/description/
 *
 * algorithms
 * Medium (28.74%)
 * Total Accepted:    196.5K
 * Total Submissions: 683.8K
 * Testcase Example:  '[2,3,-2,4]'
 *
 * Given an integer array nums, find the contiguous subarray within an array
 * (containing at least one number) which has the largest product.
 *
 * Example 1:
 *
 *
 * Input: [2,3,-2,4]
 * Output: 6
 * Explanation: [2,3] has the largest product 6.
 *
 *
 * Example 2:
 *
 *
 * Input: [-2,0,-1]
 * Output: 0
 * Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
 *
 */
func maxProduct152(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	for i := 0; i < len(nums); i++ {
		t := nums[i]
		if t > max {
			max = t
		}
		for j := i + 1; j < len(nums); j++ {
			t = t * nums[j]
			if t > max {
				max = t
			}
		}
	}

	return max
}

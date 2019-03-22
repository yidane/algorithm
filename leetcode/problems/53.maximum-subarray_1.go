package problems

/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 *
 * https://leetcode.com/problems/maximum-subarray/description/
 *
 * algorithms
 * Easy (42.57%)
 * Total Accepted:    452K
 * Total Submissions: 1.1M
 * Testcase Example:  '[-2,1,-3,4,-1,2,1,-5,4]'
 *
 * Given an integer array nums, find the contiguous subarray (containing at
 * least one number) which has the largest sum and return its sum.
 *
 * Example:
 *
 *
 * Input: [-2,1,-3,4,-1,2,1,-5,4],
 * Output: 6
 * Explanation: [4,-1,2,1] has the largest sum = 6.
 *
 *
 * Follow up:
 *
 * If you have figured out the O(n) solution, try coding another solution using
 * the divide and conquer approach, which is more subtle.
 *
 */
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	total := nums[0]
	curSum := total
	for i := 1; i < len(nums); i++ {
		curSum = curSum + nums[i]
		//-2, 3, -1, 5, 1, -5, 4
		//如果当前和小于当前元素，则当前最大值为当前元素之后的数据，最差情况也会是当前元素
		if curSum < nums[i] {
			curSum = nums[i]
		}

		if total < curSum {
			total = curSum
		}
	}

	return total
}

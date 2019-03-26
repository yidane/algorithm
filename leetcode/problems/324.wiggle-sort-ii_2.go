package problems

import "sort"

/*
 * @lc app=leetcode id=324 lang=golang
 *
 * [324] Wiggle Sort II
 *
 * https://leetcode.com/problems/wiggle-sort-ii/description/
 *
 * algorithms
 * Medium (27.56%)
 * Total Accepted:    55.5K
 * Total Submissions: 201.2K
 * Testcase Example:  '[1,5,1,1,6,4]'
 *
 * Given an unsorted array nums, reorder it such that nums[0] < nums[1] >
 * nums[2] < nums[3]....
 *
 * Example 1:
 *
 *
 * Input: nums = [1, 5, 1, 1, 6, 4]
 * Output: One possible answer is [1, 4, 1, 5, 1, 6].
 *
 * Example 2:
 *
 *
 * Input: nums = [1, 3, 2, 2, 3, 1]
 * Output: One possible answer is [2, 3, 1, 3, 1, 2].
 *
 * Note:
 * You may assume all input has valid answer.
 *
 * Follow Up:
 * Can you do it in O(n) time and/or in-place with O(1) extra space?
 */
//todo 未实现
func wiggleSort(nums []int) {
	sort.Ints(nums)

	i, j := 1, len(nums)-1

	for i < j && j < len(nums) {
		nums[i], nums[j] = nums[j], nums[i]
		i += 2
		j -= 2
	}
}

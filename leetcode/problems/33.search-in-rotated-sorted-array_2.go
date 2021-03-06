package problems

/*
 * @lc app=leetcode id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 *
 * https://leetcode.com/problems/search-in-rotated-sorted-array/description/
 *
 * algorithms
 * Medium (32.78%)
 * Total Accepted:    396.1K
 * Total Submissions: 1.2M
 * Testcase Example:  '[4,5,6,7,0,1,2]\n0'
 *
 * Suppose an array sorted in ascending order is rotated at some pivot unknown
 * to you beforehand.
 *
 * (i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).
 *
 * You are given a target value to search. If found in the array return its
 * index, otherwise return -1.
 *
 * You may assume no duplicate exists in the array.
 *
 * Your algorithm's runtime complexity must be in the order of O(log n).
 *
 * Example 1:
 *
 *
 * Input: nums = [4,5,6,7,0,1,2], target = 0
 * Output: 4
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [4,5,6,7,0,1,2], target = 3
 * Output: -1
 *
 */

//o(log n),二分法
func search33(nums []int, target int) int {
	l := len(nums)
	for i, j := 0, l-1; i <= j; {
		left := nums[i]
		right := nums[j]

		m := (j + i) / 2

		mid := nums[m]
		if mid == target {
			return m
		}

		if mid < right {
			if mid < target && right >= target {
				i = m + 1
			} else {
				j = m - 1
			}
		} else {
			if left <= target && mid > target {
				j = m - 1
			} else {
				i = m + 1
			}
		}
	}

	return -1
}

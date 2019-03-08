package problems

/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 *
 * https://leetcode.com/problems/majority-element/description/
 *
 * algorithms
 * Easy (51.44%)
 * Total Accepted:    351.5K
 * Total Submissions: 681.4K
 * Testcase Example:  '[3,2,3]'
 *
 * Given an array of size n, find the majority element. The majority element is
 * the element that appears more than ⌊ n/2 ⌋ times.
 *
 * You may assume that the array is non-empty and the majority element always
 * exist in the array.
 *
 * Example 1:
 *
 *
 * Input: [3,2,3]
 * Output: 3
 *
 * Example 2:
 *
 *
 * Input: [2,2,1,1,1,2,2]
 * Output: 2
 *
 *
 */
func majorityElement(nums []int) int {
	var m = make(map[int]int)
	maj := len(nums) / 2

	for _, i := range nums {
		v := m[i]
		v++
		if v > maj {
			return i
		}

		m[i] = v
	}

	return -1
}

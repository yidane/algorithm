package problems

import (
	"sort"
)

/*
 * @lc app=leetcode id=16 lang=golang
 *
 * [16] 3Sum Closest
 *
 * https://leetcode.com/problems/3sum-closest/description/
 *
 * algorithms
 * Medium (41.95%)
 * Total Accepted:    305.6K
 * Total Submissions: 724.4K
 * Testcase Example:  '[-1,2,1,-4]\n1'
 *
 * Given an array nums of n integers and an integer target, find three integers
 * in nums such that the sum is closest to target. Return the sum of the three
 * integers. You may assume that each input would have exactly one solution.
 *
 * Example:
 *
 *
 * Given array nums = [-1, 2, 1, -4], and target = 1.
 *
 * The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
 *
 *
 */

func threeSumClosest(array []int, target int) int {
	sub := 1 << 32
	t := 0
	sort.Ints(array)
	for k, v := range array {
		i := 0
		j := len(array) - 1
		for i < k && j > k {
			total := array[i] + array[j] + v
			switch {
			case total < target:
				i++
			case total == target:
				return total
			case total > target:
				j--
			}
			s := total - target
			if s < 0 {
				s = -s
			}
			if s < sub {
				t = total
				sub = s
			}
		}
	}

	return t
}

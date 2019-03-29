package problems

/*
 * @lc app=leetcode id=414 lang=golang
 *
 * [414] Third Maximum Number
 *
 * https://leetcode.com/problems/third-maximum-number/description/
 *
 * algorithms
 * Easy (28.78%)
 * Total Accepted:    87K
 * Total Submissions: 302.2K
 * Testcase Example:  '[3,2,1]'
 *
 * Given a non-empty array of integers, return the third maximum number in this
 * array. If it does not exist, return the maximum number. The time complexity
 * must be in O(n).
 *
 * Example 1:
 *
 * Input: [3, 2, 1]
 *
 * Output: 1
 *
 * Explanation: The third maximum is 1.
 *
 *
 *
 * Example 2:
 *
 * Input: [1, 2]
 *
 * Output: 2
 *
 * Explanation: The third maximum does not exist, so the maximum (2) is
 * returned instead.
 *
 *
 *
 * Example 3:
 *
 * Input: [2, 2, 3, 1]
 *
 * Output: 1
 *
 * Explanation: Note that the third maximum here means the third maximum
 * distinct number.
 * Both numbers with value 2 are both considered as second maximum.
 *
 *
 */
func thirdMax(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		} else {
			return nums[1]
		}
	}

	m1, m2, m3 := nums[0], 0, 0

	count := 1

	for i := 1; i < len(nums); i++ {
		v := nums[i]

		switch count {
		case 3:
			if v == m3 {
				continue
			}
			fallthrough
		case 2:
			if v == m2 {
				continue
			}
			fallthrough
		case 1:
			if v == m1 {
				continue
			}
		}

		if count == 1 {
			if v > m1 {
				m2 = m1
				m1 = v
			} else {
				m2 = v
			}
			count = 2
			continue
		}

		if count == 2 {
			if v > m1 {
				m3 = m2
				m2 = m1
				m1 = v
			} else {
				if v > m2 {
					m3 = m2
					m2 = v
				} else {
					m3 = v
				}
			}
			count = 3
			continue
		}

		if count == 3 {
			if v > m1 {
				m3 = m2
				m2 = m1
				m1 = v
			} else {
				if v > m2 {
					m3 = m2
					m2 = v
				} else {
					if v > m3 {
						m3 = v
					}
				}
			}
		}
	}

	if count == 3 {
		return m3
	} else {
		return m1
	}
}

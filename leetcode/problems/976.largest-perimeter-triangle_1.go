package problems

import "sort"

/*
 * @lc app=leetcode id=976 lang=golang
 *
 * [976] Largest Perimeter Triangle
 *
 * https://leetcode.com/problems/largest-perimeter-triangle/description/
 *
 * algorithms
 * Easy (57.06%)
 * Total Accepted:    11.8K
 * Total Submissions: 20.8K
 * Testcase Example:  '[2,1,2]'
 *
 * Given an array A of positive lengths, return the largest perimeter of a
 * triangle with non-zero area, formed from 3 of these lengths.
 *
 * If it is impossible to form any triangle of non-zero area, return 0.
 *
 *
 *
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [2,1,2]
 * Output: 5
 *
 *
 *
 * Example 2:
 *
 *
 * Input: [1,2,1]
 * Output: 0
 *
 *
 *
 * Example 3:
 *
 *
 * Input: [3,2,3,4]
 * Output: 10
 *
 *
 *
 * Example 4:
 *
 *
 * Input: [3,6,2,3]
 * Output: 8
 *
 *
 *
 *
 * Note:
 *
 *
 * 3 <= A.length <= 10000
 * 1 <= A[i] <= 10^6
 *
 *
 *
 *
 *
 */
func largestPerimeter(A []int) int {
	sort.Ints(A)
	l := len(A)

	var ai, bi, ci = 3, 2, 1
	var a, b, c int
	for {
		if ai > l || bi > l || ci > l {
			return 0
		}

		c = A[l-ci]
		b = A[l-bi]
		a = A[l-ai]

		if c >= a+b {
			if ai < bi {
				ci = bi + 1
			} else {
				ci = ai + 1
			}
			continue
		}

		if a >= c+b {
			if ci < bi {
				ai = bi + 1
			} else {
				ai = ci + 1
			}
			continue
		}

		if b >= a+c {
			if ai < ci {
				bi = ci + 1
			} else {
				bi = ai + 1
			}
			continue
		}

		break
	}

	return A[l-ai] + A[l-bi] + A[l-ci]
}

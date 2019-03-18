package problems

import "sort"

/*
 * @lc app=leetcode id=1005 lang=golang
 *
 * [1005] Maximize Sum Of Array After K Negations
 *
 * https://leetcode.com/problems/maximize-sum-of-array-after-k-negations/description/
 *
 * algorithms
 * Easy (48.98%)
 * Total Accepted:    6.6K
 * Total Submissions: 13.6K
 * Testcase Example:  '[4,2,3]\n1'
 *
 * Given an array A of integers, we must modify the array in the following way:
 * we choose an i and replace A[i] with -A[i], and we repeat this process K
 * times in total.  (We may choose the same index i multiple times.)
 *
 * Return the largest possible sum of the array after modifying it in this
 * way.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: A = [4,2,3], K = 1
 * Output: 5
 * Explanation: Choose indices (1,) and A becomes [4,-2,3].
 *
 *
 *
 * Example 2:
 *
 *
 * Input: A = [3,-1,0,2], K = 3
 * Output: 6
 * Explanation: Choose indices (1, 2, 2) and A becomes [3,1,0,2].
 *
 *
 *
 * Example 3:
 *
 *
 * Input: A = [2,-3,-1,5,-4], K = 2
 * Output: 13
 * Explanation: Choose indices (1, 4) and A becomes [2,3,-1,5,4].
 *
 *
 *
 *
 *
 *
 * Note:
 *
 *
 * 1 <= A.length <= 10000
 * 1 <= K <= 10000
 * -100 <= A[i] <= 100
 *
 *
 */

//每次翻转最小值
// 可以考虑记录几次索引，如果这几次索引值所表示的最小值相同，则不需要继续循环了。
func largestSumAfterKNegations(A []int, K int) int {
	for K > 0 {
		min := 0
		for i := 1; i < len(A); i++ {
			if A[i] < A[min] {
				min = i
			}
		}

		A[min] = -A[min]
		K--
	}

	sum := 0
	for i := 0; i < len(A); i++ {
		sum += A[i]
	}

	return sum
}

func largestSumAfterKNegations1(A []int, K int) int {
	if len(A) == 0 {
		return 0
	}

	sort.Ints(A)

	for A[0] != 0 && K > 0 {
		A[0] = -A[0]
		i := 0
		for i = len(A) - 1; i > 0; i-- {
			if A[i] < A[0] {
				A[i], A[0] = A[0], A[i]
			}
		}

		K--
	}

	sum := 0
	for i := 0; i < len(A); i++ {
		sum += A[i]
	}

	return sum
}

//todo：没明白
func largestSumAfterKNegations2(A []int, K int) int {
	sum := sum(A)
	c := count(A)
	currIndex := 0
	for K > 0 {
		if c[currIndex] != 0 {
			if currIndex < 100 {
				// negative case
				if c[currIndex] >= K {
					sum += K * 2 * (100 - currIndex)
					K = 0
				} else {
					K = K - c[currIndex]
					sum += c[currIndex] * 2 * (100 - currIndex)
					c[200-currIndex] += c[currIndex]
				}
			} else {
				// positive case
				if K%2 == 0 {
					K = 0
				} else {
					K = 0
					sum -= 2 * (currIndex - 100)
				}
			}
		}
		currIndex++
	}
	return sum
}

func count(A []int) [201]int {
	var c [201]int
	for _, v := range A {
		c[v+100]++
	}
	return c
}

func sum(A []int) int {
	sum := 0
	for _, v := range A {
		sum += v
	}
	return sum
}

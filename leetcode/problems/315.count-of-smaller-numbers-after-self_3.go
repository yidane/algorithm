package problems

/*
 * @lc app=leetcode id=315 lang=golang
 *
 * [315] Count of Smaller Numbers After Self
 *
 * https://leetcode.com/problems/count-of-smaller-numbers-after-self/description/
 *
 * algorithms
 * Hard (37.70%)
 * Total Accepted:    73K
 * Total Submissions: 193.7K
 * Testcase Example:  '[5,2,6,1]'
 *
 * You are given an integer array nums and you have to return a new counts
 * array. The counts array has the property where counts[i] is the number of
 * smaller elements to the right of nums[i].
 *
 * Example:
 *
 *
 * Input: [5,2,6,1]
 * Output: [2,1,1,0]
 * Explanation:
 * To the right of 5 there are 2 smaller elements (2 and 1).
 * To the right of 2 there is only 1 smaller element (1).
 * To the right of 6 there is 1 smaller element (1).
 * To the right of 1 there is 0 smaller element.
 *
 *
 */

//n^2
func countSmaller(nums []int) []int {
	result := make([]int, len(nums))

	for i := 0; i < len(nums)-1; i++ {
		cur := nums[i]
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < cur {
				result[i] += 1
			}
		}
	}

	return result
}

//func countSmaller1(nums []int) []int {
//	return nil
//}

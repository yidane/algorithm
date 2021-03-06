package problems

/*
 * @lc app=leetcode id=198 lang=golang
 *
 * [198] House Robber
 *
 * https://leetcode.com/problems/house-robber/description/
 *
 * algorithms
 * Easy (40.90%)
 * Total Accepted:    309.3K
 * Total Submissions: 756.1K
 * Testcase Example:  '[1,2,3,1]'
 *
 * You are a professional robber planning to rob houses along a street. Each
 * house has a certain amount of money stashed, the only constraint stopping
 * you from robbing each of them is that adjacent houses have security system
 * connected and it will automatically contact the police if two adjacent
 * houses were broken into on the same night.
 *
 * Given a list of non-negative integers representing the amount of money of
 * each house, determine the maximum amount of money you can rob tonight
 * without alerting the police.
 *
 * Example 1:
 *
 *
 * Input: [1,2,3,1]
 * Output: 4
 * Explanation: Rob house 1 (money = 1) and then rob house 3 (money =
 * 3).
 * Total amount you can rob = 1 + 3 = 4.
 *
 * Example 2:
 *
 *
 * Input: [2,7,9,3,1]
 * Output: 12
 * Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house
 * 5 (money = 1).
 * Total amount you can rob = 2 + 9 + 1 = 12.
 *
 *
 */
//动态规划算法
//http://www.cnblogs.com/grandyang/p/4383632.html
func rob(nums []int) int {
	robEven := 0
	robOdd := 0
	n := len(nums)

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			newRobEven := robEven + nums[i]
			if newRobEven > robOdd {
				robEven = newRobEven
			} else {
				robEven = robOdd
			}
		} else {
			newRobOdd := robOdd + nums[i]
			if newRobOdd > robEven {
				robOdd = newRobOdd
			} else {
				robOdd = robEven
			}
		}
	}

	if robOdd > robEven {
		return robOdd
	}

	return robEven
}

//递归实现，比较好理解
//通过递归，可以推导出动态规划算法公式
func rob1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

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

	odd := nums[0] + rob(nums[2:])
	even := nums[1] + rob(nums[3:])

	if odd > even {
		return odd
	}

	return even
}

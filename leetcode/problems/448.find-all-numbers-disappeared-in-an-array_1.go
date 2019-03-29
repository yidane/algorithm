package problems

/*
 * @lc app=leetcode id=448 lang=golang
 *
 * [448] Find All Numbers Disappeared in an Array
 *
 * https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/description/
 *
 * algorithms
 * Easy (52.92%)
 * Total Accepted:    141.4K
 * Total Submissions: 267.2K
 * Testcase Example:  '[4,3,2,7,8,2,3,1]'
 *
 * Given an array of integers where 1 ≤ a[i] ≤ n (n = size of array), some
 * elements appear twice and others appear once.
 *
 * Find all the elements of [1, n] inclusive that do not appear in this array.
 *
 * Could you do it without extra space and in O(n) runtime? You may assume the
 * returned list does not count as extra space.
 *
 * Example:
 *
 * Input:
 * [4,3,2,7,8,2,3,1]
 *
 * Output:
 * [5,6]
 *
 *
 */
func findDisappearedNumbers(nums []int) []int {
	var result []int
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = 1
	}

	for i := 1; i < len(nums)+1; i++ {
		if m[i] == 0 {
			result = append(result, i)
		}
	}

	return result
}

func findDisappearedNumbers1(nums []int) []int {
	var result []int

	//结题思路
	//将数值交换到相对应的顺序位置，若原位置本就等于该值，则表示该值重复
	for i := 0; i < len(nums); i++ {
		e := nums[i]        //获取当前位置值
		if nums[e-1] != e { //若当前位置的值不等于当前位置值所对应索引位置的值
			nums[e-1], nums[i] = nums[i], nums[e-1] //交换，使得当前位置的值等于当前位置值所对应索引的值
			i--                                     //右边回撤一步，继续判断
		}
	}

	for i := 0; i < len(nums); i++ {
		if i+1 != nums[i] {
			result = append(result, i+1)
		}
	}

	return result
}

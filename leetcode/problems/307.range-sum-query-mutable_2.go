package problems

/*
 * @lc app=leetcode id=307 lang=golang
 *
 * [307] Range Sum Query - Mutable
 *
 * https://leetcode.com/problems/range-sum-query-mutable/description/
 *
 * algorithms
 * Medium (28.13%)
 * Total Accepted:    69K
 * Total Submissions: 245.3K
 * Testcase Example:  '["NumArray","sumRange","update","sumRange"]\n[[[1,3,5]],[0,2],[1,2],[0,2]]'
 *
 * Given an integer array nums, find the sum of the elements between indices i
 * and j (i ≤ j), inclusive.
 *
 * The update(i, val) function modifies nums by updating the element at index i
 * to val.
 *
 * Example:
 *
 *
 * Given nums = [1, 3, 5]
 *
 * sumRange(0, 2) -> 9
 * update(1, 2)
 * sumRange(0, 2) -> 8
 *
 *
 * Note:
 *
 *
 * The array is only modifiable by the update function.
 * You may assume the number of calls to update and sumRange function is
 * distributed evenly.
 *
 *
 */
type NumArray struct {
	nums []int
}

func NumArrayConstructor(nums []int) NumArray {
	return NumArray{nums: nums}
}

func (this *NumArray) Update(i int, val int) {
	this.nums[i] = val
}

func (this *NumArray) SumRange(i int, j int) int {
	t := 0
	for k := i; k <= j; k++ {
		t += this.nums[k]
	}

	return t
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(i,val);
 * param_2 := obj.SumRange(i,j);
 */

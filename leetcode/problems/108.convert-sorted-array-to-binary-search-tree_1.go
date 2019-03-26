package problems

/*
 * @lc app=leetcode id=108 lang=golang
 *
 * [108] Convert Sorted Array to Binary Search Tree
 *
 * https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/description/
 *
 * algorithms
 * Easy (49.50%)
 * Total Accepted:    245.2K
 * Total Submissions: 493.9K
 * Testcase Example:  '[-10,-3,0,5,9]'
 *
 * Given an array where elements are sorted in ascending order, convert it to a
 * height balanced BST.
 *
 * For this problem, a height-balanced binary tree is defined as a binary tree
 * in which the depth of the two subtrees of every node never differ by more
 * than 1.
 *
 * Example:
 *
 *
 * Given the sorted array: [-10,-3,0,5,9],
 *
 * One possible answer is: [0,-3,9,-10,null,5], which represents the following
 * height balanced BST:
 *
 * ⁠     0
 * ⁠    / \
 * ⁠  -3   9
 * ⁠  /   /
 * ⁠-10  5
 *
 *
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	return newBSTNode(nums, 0, len(nums)-1)
}

func newBSTNode(nums []int, begin, end int) *TreeNode {
	if begin > end {
		return nil
	}

	mid := (begin + end) / 2
	node := &TreeNode{Val: nums[mid]}
	node.Left = newBSTNode(nums, begin, mid-1)
	node.Right = newBSTNode(nums, mid+1, end)

	return node
}

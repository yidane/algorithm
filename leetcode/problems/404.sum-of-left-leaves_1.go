package problems

/*
 * @lc app=leetcode id=404 lang=golang
 *
 * [404] Sum of Left Leaves
 *
 * https://leetcode.com/problems/sum-of-left-leaves/description/
 *
 * algorithms
 * Easy (48.74%)
 * Total Accepted:    119K
 * Total Submissions: 244.1K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Find the sum of all left leaves in a given binary tree.
 *
 * Example:
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 * There are two left leaves in the binary tree, with values 9 and 15
 * respectively. Return 24.
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
func sumOfLeftLeaves(root *TreeNode) int {
	return sumLeft(root, false)
}

func sumLeft(node *TreeNode, contain bool) int {
	if node == nil {
		return 0
	}

	sum := 0
	if contain && node.Left == nil && node.Right == nil {
		sum += node.Val
	}

	sum += sumLeft(node.Left, true)
	sum += sumLeft(node.Right, false)

	return sum
}

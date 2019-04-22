package problems

/*
 * @lc app=leetcode id=110 lang=golang
 *
 * [110] Balanced Binary Tree
 *
 * https://leetcode.com/problems/balanced-binary-tree/description/
 *
 * algorithms
 * Easy (40.70%)
 * Total Accepted:    310.4K
 * Total Submissions: 762.3K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, determine if it is height-balanced.
 *
 * For this problem, a height-balanced binary tree is defined as:
 *
 *
 * a binary tree in which the depth of the two subtrees of every node never
 * differ by more than 1.
 *
 *
 * Example 1:
 *
 * Given the following tree [3,9,20,null,null,15,7]:
 *
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 * Return true.
 *
 * Example 2:
 *
 * Given the following tree [1,2,2,3,3,null,null,4,4]:
 *
 *
 * ⁠      1
 * ⁠     / \
 * ⁠    2   2
 * ⁠   / \
 * ⁠  3   3
 * ⁠ / \
 * ⁠4   4
 *
 *
 * Return false.
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
func isBalanced(root *TreeNode) bool {
	return checkBalanced(root) != -1
}

func checkBalanced(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := checkBalanced(root.Left)
	if left == -1 {
		return -1
	}

	right := checkBalanced(root.Right)
	if right == -1 {
		return -1
	}

	if right < left {
		left, right = right, left
	}

	if right-left > 1 {
		return -1
	}
	return right + 1
}

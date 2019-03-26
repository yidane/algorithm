package problems

/*
 * @lc app=leetcode id=543 lang=golang
 *
 * [543] Diameter of Binary Tree
 *
 * https://leetcode.com/problems/diameter-of-binary-tree/description/
 *
 * algorithms
 * Easy (46.27%)
 * Total Accepted:    117.9K
 * Total Submissions: 254.6K
 * Testcase Example:  '[1,2,3,4,5]'
 *
 *
 * Given a binary tree, you need to compute the length of the diameter of the
 * tree. The diameter of a binary tree is the length of the longest path
 * between any two nodes in a tree. This path may or may not pass through the
 * root.
 *
 *
 *
 * Example:
 * Given a binary tree
 *
 * ⁠         1
 * ⁠        / \
 * ⁠       2   3
 * ⁠      / \
 * ⁠     4   5
 *
 *
 *
 * Return 3, which is the length of the path [4,2,1,3] or [5,2,1,3].
 *
 *
 * Note:
 * The length of path between two nodes is represented by the number of edges
 * between them.
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
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxDepth := findMaxDepth(root.Left) + findMaxDepth(root.Right)

	leftMaxDepth := diameterOfBinaryTree(root.Left)
	rightMaxDepth := diameterOfBinaryTree(root.Right)

	if leftMaxDepth > maxDepth {
		maxDepth = leftMaxDepth
	}

	if rightMaxDepth > maxDepth {
		maxDepth = rightMaxDepth
	}

	return maxDepth
}

func findMaxDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}

	depthLeft := findMaxDepth(node.Left)
	depthRight := findMaxDepth(node.Right)

	if depthLeft < depthRight {
		return depthRight + 1
	} else {
		return depthLeft + 1
	}
}

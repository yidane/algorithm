package problems

/*
 * @lc app=leetcode id=111 lang=golang
 *
 * [111] Minimum Depth of Binary Tree
 *
 * https://leetcode.com/problems/minimum-depth-of-binary-tree/description/
 *
 * algorithms
 * Easy (35.10%)
 * Total Accepted:    286.5K
 * Total Submissions: 816K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, find its minimum depth.
 *
 * The minimum depth is the number of nodes along the shortest path from the
 * root node down to the nearest leaf node.
 *
 * Note: A leaf is a node with no children.
 *
 * Example:
 *
 * Given binary tree [3,9,20,null,null,15,7],
 *
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 * return its minimum depth = 2.
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
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	minDepth := 1
	var nodes []*TreeNode
	if root.Left != nil {
		nodes = append(nodes, root.Left)
	}
	if root.Right != nil {
		nodes = append(nodes, root.Right)
	}

	for len(nodes) > 0 {
		var newNodes []*TreeNode

		for i := 0; i < len(nodes); i++ {
			node := nodes[i]

			if node.Left == nil && node.Right == nil {
				return minDepth + 1
			}

			if node.Left != nil {
				newNodes = append(newNodes, node.Left)
			}

			if node.Right != nil {
				newNodes = append(newNodes, node.Right)
			}
		}

		minDepth++
		nodes = newNodes
	}

	return minDepth
}

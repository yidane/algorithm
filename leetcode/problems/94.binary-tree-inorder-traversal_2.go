package problems

/*
 * @lc app=leetcode id=94 lang=golang
 *
 * [94] Binary Tree Inorder Traversal
 *
 * https://leetcode.com/problems/binary-tree-inorder-traversal/description/
 *
 * algorithms
 * Medium (55.23%)
 * Total Accepted:    416.8K
 * Total Submissions: 754.5K
 * Testcase Example:  '[1,null,2,3]'
 *
 * Given a binary tree, return the inorder traversal of its nodes' values.
 *
 * Example:
 *
 *
 * Input: [1,null,2,3]
 * ⁠  1
 * ⁠   \
 * ⁠    2
 * ⁠   /
 * ⁠  3
 *
 * Output: [1,3,2]
 *
 * Follow up: Recursive solution is trivial, could you do it iteratively?
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
func inorderTraversal(root *TreeNode) []int {
	var result []int
	if root != nil {
		result = append(result, inorderTraversal(root.Left)...)
		result = append(result, root.Val)
		result = append(result, inorderTraversal(root.Right)...)
	}
	return result
}

func inorderTraversal1(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	if root == nil {
		return res
	}
	if root.Right != nil {
		stack = append(stack, root.Right)
		root.Right = nil
	}
	stack = append(stack, root)
	if root.Left != nil {
		stack = append(stack, root.Left)
		root.Left = nil
	}
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Left == nil {
			res = append(res, node.Val)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
			node.Right = nil
		}
		if node.Left != nil {
			stack = append(stack, node)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
			node.Left = nil
		}
	}
	return res
}

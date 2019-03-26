package problems

/*
 * @lc app=leetcode id=701 lang=golang
 *
 * [701] Insert into a Binary Search Tree
 *
 * https://leetcode.com/problems/insert-into-a-binary-search-tree/description/
 *
 * algorithms
 * Medium (74.22%)
 * Total Accepted:    34.8K
 * Total Submissions: 46.8K
 * Testcase Example:  '[4,2,7,1,3]\n5'
 *
 * Given the root node of a binary search tree (BST) and a value to be inserted
 * into the tree, insert the value into the BST. Return the root node of the
 * BST after the insertion. It is guaranteed that the new value does not exist
 * in the original BST.
 *
 * Note that there may exist multiple valid ways for the insertion, as long as
 * the tree remains a BST after insertion. You can return any of them.
 *
 * For example,
 *
 *
 * Given the tree:
 * ⁠       4
 * ⁠      / \
 * ⁠     2   7
 * ⁠    / \
 * ⁠   1   3
 * And the value to insert: 5
 *
 *
 * You can return this binary search tree:
 *
 *
 * ⁠        4
 * ⁠      /   \
 * ⁠     2     7
 * ⁠    / \   /
 * ⁠   1   3 5
 *
 *
 * This tree is also valid:
 *
 *
 * ⁠        5
 * ⁠      /   \
 * ⁠     2     7
 * ⁠    / \
 * ⁠   1   3
 * ⁠        \
 * ⁠         4
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
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	if val < root.Val {
		if root.Left == nil {
			root.Left = &TreeNode{Val: val}
		} else {
			insertIntoBST(root.Left, val)
		}
	} else {

		if root.Right == nil {
			root.Right = &TreeNode{Val: val}
		} else {
			insertIntoBST(root.Right, val)
		}
	}

	return root
}

func insertIntoBST1(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	originalRoot := root

	for {
		if val < root.Val {
			if root.Left == nil {
				root.Left = &TreeNode{Val: val}
				break
			} else {
				root = root.Left
			}
		} else {
			if root.Right == nil {
				root.Right = &TreeNode{Val: val}
				break
			} else {
				root = root.Right
			}
		}
	}

	return originalRoot
}

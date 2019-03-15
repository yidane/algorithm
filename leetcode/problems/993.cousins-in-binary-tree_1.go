package problems

/*
 * @lc app=leetcode id=993 lang=golang
 *
 * [993] Cousins in Binary Tree
 *
 * https://leetcode.com/problems/cousins-in-binary-tree/description/
 *
 * algorithms
 * Easy (53.05%)
 * Total Accepted:    8.5K
 * Total Submissions: 15.9K
 * Testcase Example:  '[1,2,3,4]\n4\n3'
 *
 * In a binary tree, the root node is at depth 0, and children of each depth k
 * node are at depth k+1.
 *
 * Two nodes of a binary tree are cousins if they have the same depth, but have
 * different parents.
 *
 * We are given the root of a binary tree with unique values, and the values x
 * and y of two different nodes in the tree.
 *
 * Return true if and only if the nodes corresponding to the values x and y are
 * cousins.
 *
 *
 *
 * Example 1:
 *
 *
 *
 * Input: root = [1,2,3,4], x = 4, y = 3
 * Output: false
 *
 *
 *
 * Example 2:
 *
 *
 *
 * Input: root = [1,2,3,null,4,null,5], x = 5, y = 4
 * Output: true
 *
 *
 *
 * Example 3:
 *
 *
 *
 *
 * Input: root = [1,2,3,null,4], x = 2, y = 3
 * Output: false
 *
 *
 *
 *
 *
 * Note:
 *
 *
 * The number of nodes in the tree will be between 2 and 100.
 * Each node has a unique integer value from 1 to 100.
 *
 *
 *
 *
 *
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
func isCousins(root *TreeNode, x int, y int) bool {
	//第一层节点不可能为堂兄弟
	if x == y || root == nil || root.Val == x || root.Val == y {
		return false
	}

	//若不存在第二层，则不存在堂兄弟
	if root.Left == nil || root.Right == nil {
		return false
	}

	//第二层的节点不存在堂兄弟
	if root.Left.Val == x || root.Left.Val == y || root.Right.Val == x || root.Right.Val == y {
		return false
	}

	px, kx, cx := findLocation(root.Left, 0, x)
	py, ky, cy := findLocation(root.Right, 0, y)

	return cx && cy && kx == ky && px != py
}

func findLocation(node *TreeNode, l, x int) (parent, k int, contain bool) {
	root := node

	if root.Val > x {
		if root.Left == nil {
			return -1, -1, false
		}

		if root.Left.Val == x {
			return root.Val, l + 1, true
		}

		return findLocation(root.Left, l+1, x)
	} else {
		if root.Right == nil {
			return -1, -1, false
		}

		if root.Right.Val == x {
			return root.Val, l + 1, true
		}

		return findLocation(root.Right, l+1, x)
	}
}

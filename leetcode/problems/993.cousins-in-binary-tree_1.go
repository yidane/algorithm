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

type n struct {
	parent int
	k      int
	v      int
}

func (n *n) hasValue() bool {
	return n.parent != 0
}

func (n *n) setValue(p, k int) {
	n.parent = p
	n.k = k
}

func (n *n) equal(an *n) bool {
	if n == nil || an == nil {
		return false
	}

	return n.parent != an.parent && n.k == an.k
}

func isCousins(root *TreeNode, x int, y int) bool {
	xn := &n{v: x}
	yn := &n{v: y}

	findLocation(root, -1, 0, xn, yn)

	return xn.equal(yn)
}

func findLocation(root *TreeNode, p, k int, xn, yn *n) {
	if root == nil {
		return
	}

	if root.Val == xn.v {
		xn.setValue(p, k)
	}

	if root.Val == yn.v {
		yn.setValue(p, k)
	}

	if xn.hasValue() && yn.hasValue() {
		return
	}

	findLocation(root.Left, root.Val, k+1, xn, yn)
	findLocation(root.Right, root.Val, k+1, xn, yn)
}

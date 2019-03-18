package problems

import "container/list"

/*
 * @lc app=leetcode id=653 lang=golang
 *
 * [653] Two Sum IV - Input is a BST
 *
 * https://leetcode.com/problems/two-sum-iv-input-is-a-bst/description/
 *
 * algorithms
 * Easy (52.02%)
 * Total Accepted:    78.2K
 * Total Submissions: 150.3K
 * Testcase Example:  '[5,3,6,2,4,null,7]\n9'
 *
 * Given a Binary Search Tree and a target number, return true if there exist
 * two elements in the BST such that their sum is equal to the given target.
 *
 * Example 1:
 *
 *
 * Input:
 * ⁠   5
 * ⁠  / \
 * ⁠ 3   6
 * ⁠/ \   \
 * 2   4   7
 *
 * Target = 9
 *
 * Output: True
 *
 *
 *
 *
 * Example 2:
 *
 *
 * Input:
 * ⁠   5
 * ⁠  / \
 * ⁠ 3   6
 * ⁠/ \   \
 * 2   4   7
 *
 * Target = 28
 *
 * Output: False
 *
 *
 *
 * @lc app=leetcode id=653 lang=golang
 *
 * [653] Two Sum IV - Input is a BST
 *
 * https://leetcode.com/problems/two-sum-iv-input-is-a-bst/description/
 *
 * algorithms
 * Easy (52.02%)
 * Total Accepted:    78.2K
 * Total Submissions: 150.3K
 * Testcase Example:  '[5,3,6,2,4,null,7]\n9'
 *
 * Given a Binary Search Tree and a target number, return true if there exist
 * two elements in the BST such that their sum is equal to the given target.
 *
 * Example 1:
 *
 *
 * Input:
 * ⁠   5
 * ⁠  / \
 * ⁠ 3   6
 * ⁠/ \   \
 * 2   4   7
 *
 * Target = 9
 *
 * Output: True
 *
 *
 *
 *
 * Example 2:
 *
 *
 * Input:
 * ⁠   5
 * ⁠  / \
 * ⁠ 3   6
 * ⁠/ \   \
 * 2   4   7
 *
 * Target = 28
 *
 * Output: False
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

func findTarget(root *TreeNode, k int) bool {
	var l = list.New()

	prev(root, l)

	head, tail := l.Front(), l.Back()

	for head != tail {
		s := head.Value.(int) + tail.Value.(int)
		if s == k {
			return true
		} else if s < k {
			head = head.Next()
		} else {
			tail = tail.Prev()
		}
	}

	return false
}

func prev(root *TreeNode, m *list.List) {
	if root == nil {
		return
	}

	prev(root.Left, m)

	m.PushBack(root.Val)

	prev(root.Right, m)
}

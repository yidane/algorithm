package problems

/*
 * @lc app=leetcode id=107 lang=golang
 *
 * [107] Binary Tree Level Order Traversal II
 *
 * https://leetcode.com/problems/binary-tree-level-order-traversal-ii/description/
 *
 * algorithms
 * Easy (46.26%)
 * Total Accepted:    219.2K
 * Total Submissions: 473.4K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, return the bottom-up level order traversal of its
 * nodes' values. (ie, from left to right, level by level from leaf to root).
 *
 *
 * For example:
 * Given binary tree [3,9,20,null,null,15,7],
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 *
 *
 * return its bottom-up level order traversal as:
 *
 * [
 * ⁠ [15,7],
 * ⁠ [9,20],
 * ⁠ [3]
 * ]
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
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	arr := [][]int{{root.Val}}

	nodes := []*TreeNode{root.Left, root.Right}

	for len(nodes) > 0 {
		var newNodes []*TreeNode
		var newArr []int
		for i := 0; i < len(nodes); i++ {
			node := nodes[i]
			if node != nil {
				newArr = append(newArr, node.Val)
				newNodes = append(newNodes, node.Left, node.Right)
			}
		}

		if len(newArr) > 0 {
			arr = append(arr, newArr)
		}
		nodes = newNodes
	}

	//此法甚妙
	length := len(arr)
	for i := 0; i < length/2; i++ {
		arr[i], arr[length-i-1] = arr[length-i-1], arr[i]
	}

	return arr
}

//此方法快一点
func levelOrderBottom1(root *TreeNode) [][]int {
	result := [][]int{}
	queue := [][]*TreeNode{}
	if root != nil {
		queue = append(queue, []*TreeNode{root})
	}
	for len(queue) != 0 && len(queue[0]) != 0 {
		nodes := queue[0]
		queue = queue[1:]
		vals := []int{}
		ns := []*TreeNode{}
		for _, node := range nodes {
			vals = append(vals, node.Val)
			if node.Left != nil {
				ns = append(ns, node.Left)
			}
			if node.Right != nil {
				ns = append(ns, node.Right)
			}
		}
		result = append(result, vals)
		queue = append(queue, ns)
	}
	length := len(result)
	for i := 0; i < length/2; i++ {
		result[i], result[length-i-1] = result[length-i-1], result[i]
	}
	return result
}

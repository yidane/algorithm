package problems

/*
 * @lc app=leetcode id=637 lang=golang
 *
 * [637] Average of Levels in Binary Tree
 *
 * https://leetcode.com/problems/average-of-levels-in-binary-tree/description/
 *
 * algorithms
 * Easy (58.12%)
 * Total Accepted:    72.2K
 * Total Submissions: 124.3K
 * Testcase Example:  '[3,9,20,15,7]'
 *
 * Given a non-empty binary tree, return the average value of the nodes on each
 * level in the form of an array.
 *
 * Example 1:
 *
 * Input:
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 * Output: [3, 14.5, 11]
 * Explanation:
 * The average value of nodes on level 0 is 3,  on level 1 is 14.5, and on
 * level 2 is 11. Hence return [3, 14.5, 11].
 *
 *
 *
 * Note:
 *
 * The range of node's value is in the range of 32-bit signed integer.
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

type lvl struct {
	total int
	l     int
}

func (l *lvl) add(v int) {
	l.total += v
	l.l++
}

func (l *lvl) avg() float64 {
	return float64(l.total) / float64(l.l)
}

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}

	var arr []*lvl
	readNodeValue(root, 0, &arr)

	result := make([]float64, len(arr))

	for i := 0; i < len(arr); i++ {
		result[i] = arr[i].avg()
	}

	return result
}

func readNodeValue(node *TreeNode, l int, arr *[]*lvl) {
	if node == nil {
		return
	}

	if len(*arr) <= l {
		*arr = append(*arr, &lvl{})
	}

	(*arr)[l].add(node.Val)

	readNodeValue(node.Left, l+1, arr)
	readNodeValue(node.Right, l+1, arr)
}

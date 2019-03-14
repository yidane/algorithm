package problems

/*
 * @lc app=leetcode id=897 lang=golang
 *
 * [897] Increasing Order Search Tree
 *
 * https://leetcode.com/problems/increasing-order-search-tree/description/
 *
 * algorithms
 * Easy (63.25%)
 * Total Accepted:    23.2K
 * Total Submissions: 36.6K
 * Testcase Example:  '[5,3,6,2,4,null,8,1,null,null,null,7,9]'
 *
 * Given a tree, rearrange the tree in in-order so that the leftmost node in
 * the tree is now the root of the tree, and every node has no left child and
 * only 1 right child.
 *
 *
 * Example 1:
 * Input: [5,3,6,2,4,null,8,1,null,null,null,7,9]
 *
 * ⁠      5
 * ⁠     / \
 * ⁠   3    6
 * ⁠  / \    \
 * ⁠ 2   4    8
 * /        / \
 * 1        7   9
 *
 * Output: [1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9]
 *
 * ⁠1
 * \
 * 2
 * \
 * 3
 * \
 * 4
 * \
 * 5
 * \
 * 6
 * \
 * 7
 * \
 * 8
 * \
 * ⁠                9
 *
 * Note:
 *
 *
 * The number of nodes in the given tree will be between 1 and 100.
 * Each node will have a unique integer value from 0 to 1000.
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
func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	arr := scanTree(root)

	result := &TreeNode{
		Val: arr[0],
	}
	resultRoot := result

	for i := 1; i < len(arr); i++ {
		child := TreeNode{
			Val: arr[i],
		}

		result.Right = &child
		result = result.Right
	}

	return resultRoot
}

func scanTree(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var arr []int

	leftArr := scanTree(root.Left)
	rightArr := scanTree(root.Right)

	if leftArr != nil {
		arr = append(arr, leftArr...)
	}

	arr = append(arr, root.Val)

	if rightArr != nil {
		arr = append(arr, rightArr...)
	}

	return arr
}

package problems

/*
 * @lc app=leetcode id=872 lang=golang
 *
 * [872] Leaf-Similar Trees
 *
 * https://leetcode.com/problems/leaf-similar-trees/description/
 *
 * algorithms
 * Easy (62.28%)
 * Total Accepted:    36.3K
 * Total Submissions: 58.3K
 * Testcase Example:  '[3,5,1,6,2,9,8,null,null,7,4]\n[3,5,1,6,7,4,2,null,null,null,null,null,null,9,8]'
 *
 * Consider all the leaves of a binary tree.  From left to right order, the
 * values of those leaves form a leaf value sequence.
 *
 *
 *
 * For example, in the given tree above, the leaf value sequence is (6, 7, 4,
 * 9, 8).
 *
 * Two binary trees are considered leaf-similar if their leaf value sequence is
 * the same.
 *
 * Return true if and only if the two given trees with head nodes root1 and
 * root2 are leaf-similar.
 *
 *
 *
 * Note:
 *
 *
 * Both of the given trees will have between 1 and 100 nodes.
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
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leftLeafArr := scanLeaf(root1)
	rightLeafArr := scanLeaf(root2)

	if len(leftLeafArr) != len(rightLeafArr) {
		return false
	}

	for i := 0; i < len(leftLeafArr); i++ {
		if leftLeafArr[i] != rightLeafArr[i] {
			return false
		}
	}

	return true
}

func isLeaf(node *TreeNode) bool {
	return node != nil && node.Left == nil && node.Right == nil
}

func scanLeaf(node *TreeNode) []int {
	if node == nil {
		return nil
	}

	var arr []int

	if isLeaf(node) {
		arr = append(arr, node.Val)
	} else {
		leftArr := scanLeaf(node.Left)
		rightArr := scanLeaf(node.Right)

		arr = append(arr, leftArr...)
		arr = append(arr, rightArr...)
	}

	return arr
}

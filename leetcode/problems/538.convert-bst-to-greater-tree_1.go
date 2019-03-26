package problems

/*
 * @lc app=leetcode id=538 lang=golang
 *
 * [538] Convert BST to Greater Tree
 *
 * https://leetcode.com/problems/convert-bst-to-greater-tree/description/
 *
 * algorithms
 * Easy (50.13%)
 * Total Accepted:    71.7K
 * Total Submissions: 142.9K
 * Testcase Example:  '[5,2,13]'
 *
 * Given a Binary Search Tree (BST), convert it to a Greater Tree such that
 * every key of the original BST is changed to the original key plus sum of all
 * keys greater than the original key in BST.
 *
 *
 * Example:
 *
 * Input: The root of a Binary Search Tree like this:
 * ⁠             5
 * ⁠           /   \
 * ⁠          2     13
 *
 * Output: The root of a Greater Tree like this:
 * ⁠            18
 * ⁠           /   \
 * ⁠         20     13
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

//尚待优化
func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	return convertBSTNode(root, root)
}

func convertBSTNode(root *TreeNode, node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}

	newNode := &TreeNode{Val: sumBST(root, node.Val)}

	newNode.Left = convertBSTNode(root, node.Left)
	newNode.Right = convertBSTNode(root, node.Right)

	return newNode
}

func sumBST(root *TreeNode, v int) int {
	if root == nil {
		return 0
	}

	sum := 0

	if root.Val < v {
		sum += sumBST(root.Right, v)
		return sum
	}

	if root.Val >= v {
		sum += root.Val
	}

	sum += sumBST(root.Right, v)
	sum += sumBST(root.Left, v)

	return sum
}

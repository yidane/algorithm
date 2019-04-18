package problems

/*
 * @lc app=leetcode id=1026 lang=golang
 *
 * [1026] Maximum Difference Between Node and Ancestor
 *
 * https://leetcode.com/problems/maximum-difference-between-node-and-ancestor/description/
 *
 * algorithms
 * Medium (57.30%)
 * Total Accepted:    4.7K
 * Total Submissions: 8.1K
 * Testcase Example:  '[8,3,10,1,6,null,14,null,null,4,7,13]'
 *
 * Given the root of a binary tree, find the maximum value V for which there
 * exists different nodes A and B where V = |A.val - B.val| and A is an
 * ancestor of B.
 *
 * (A node A is an ancestor of B if either: any child of A is equal to B, or
 * any child of A is an ancestor of B.)
 *
 *
 *
 * Example 1:
 *
 *
 *
 *
 * Input: [8,3,10,1,6,null,14,null,null,4,7,13]
 * Output: 7
 * Explanation:
 * We have various ancestor-node differences, some of which are given below :
 * |8 - 3| = 5
 * |3 - 7| = 4
 * |8 - 1| = 7
 * |10 - 13| = 3
 * Among all possible differences, the maximum value of 7 is obtained by |8 -
 * 1| = 7.
 *
 *
 *
 *
 * Note:
 *
 *
 * The number of nodes in the tree is between 2 and 5000.
 * Each node will have value between 0 and 100000.
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

//TODO:worst
func maxAncestorDiff(root *TreeNode) int {
	nodes := []*TreeNode{root}
	maxAncestor := 0

	for len(nodes) > 0 {
		var tNodes []*TreeNode
		for _, node := range nodes {
			if node.Left != nil {
				t := findMaxAncestor(node.Val, node.Left)

				if t > maxAncestor {
					maxAncestor = t
				}

				tNodes = append(tNodes, node.Left)
			}
			if node.Right != nil {
				t := findMaxAncestor(node.Val, node.Right)

				if t > maxAncestor {
					maxAncestor = t
				}

				tNodes = append(tNodes, node.Right)
			}
		}

		nodes = tNodes
	}

	return maxAncestor
}

func findMaxAncestor(v int, root *TreeNode) int {
	nodes := []*TreeNode{root}
	maxAncestor := v - root.Val
	if maxAncestor < 0 {
		maxAncestor = -maxAncestor
	}

	for len(nodes) > 0 {
		var tNodes []*TreeNode
		for _, node := range nodes {
			if node.Left != nil {
				t := node.Left.Val - v
				if t < 0 {
					t = -t
				}
				if t > maxAncestor {
					maxAncestor = t
				}

				tNodes = append(tNodes, node.Left)
			}

			if node.Right != nil {
				t := node.Right.Val - v
				if t < 0 {
					t = -t
				}
				if t > maxAncestor {
					maxAncestor = t
				}
				tNodes = append(tNodes, node.Right)
			}
		}

		nodes = tNodes
	}

	return maxAncestor
}

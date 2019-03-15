package problems

/*
 * @lc app=leetcode id=669 lang=golang
 *
 * [669] Trim a Binary Search Tree
 *
 * https://leetcode.com/problems/trim-a-binary-search-tree/description/
 *
 * algorithms
 * Easy (59.82%)
 * Total Accepted:    59.5K
 * Total Submissions: 99.5K
 * Testcase Example:  '[1,0,2]\n1\n2'
 *
 *
 * Given a binary search tree and the lowest and highest boundaries as L and R,
 * trim the tree so that all its elements lies in [L, R] (R >= L). You might
 * need to change the root of the tree, so the result should return the new
 * root of the trimmed binary search tree.
 *
 *
 * Example 1:
 *
 * Input:
 * ⁠   1
 * ⁠  / \
 * ⁠ 0   2
 *
 * ⁠ L = 1
 * ⁠ R = 2
 *
 * Output:
 * ⁠   1
 * ⁠     \
 * ⁠      2
 *
 *
 *
 * Example 2:
 *
 * Input:
 * ⁠   3
 * ⁠  / \
 * ⁠ 0   4
 * ⁠  \
 * ⁠   2
 * ⁠  /
 * ⁠ 1
 *
 * ⁠ L = 1
 * ⁠ R = 3
 *
 * Output:
 * ⁠     3
 * ⁠    /
 * ⁠  2
 * ⁠ /
 * ⁠1
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
func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root == nil {
		return nil
	}

	for {
		if root == nil {
			return nil
		}

		if root.Val < L {
			root = root.Right
			continue
		}

		if root.Val > R {
			root = root.Left
			continue
		}

		break
	}

	newRoot := &TreeNode{Val: root.Val}
	leftNode := trimBST(root.Left, L, R)
	if leftNode != nil {
		if newRoot.Val < leftNode.Val {
			newRoot.Right = leftNode
		} else {
			newRoot.Left = leftNode
		}
	}

	rightNode := trimBST(root.Right, L, R)
	if rightNode != nil {
		if newRoot.Val < rightNode.Val {
			newRoot.Right = rightNode
		} else {
			newRoot.Left = rightNode
		}
	}

	return newRoot
}

//nice
func trimBST1(root *TreeNode, L int, R int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < L {
		return trimBST(root.Right, L, R)
	}
	if root.Val > R {
		return trimBST(root.Left, L, R)
	}
	return &TreeNode{root.Val, trimBST(root.Left, L, R), trimBST(root.Right, L, R)}
}

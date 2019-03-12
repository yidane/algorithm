package problems

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(v int) *TreeNode {
	node := new(TreeNode)
	node.Val = v
	return node
}

func (t *TreeNode) AddLeft(v int) *TreeNode {
	node := NewTreeNode(v)
	t.Left = node
	return t
}

func (t *TreeNode) AddRight(v int) *TreeNode {
	node := NewTreeNode(v)
	t.Right = node
	return t
}

//func NewTreeNode(arr []int) *TreeNode {
//	if arr == nil || len(arr) == 0 {
//		return nil
//	}
//
//	t := new(TreeNode)
//	t.Val = arr[0]
//
//	for i := 1; i < len(arr); i++ {
//
//	}
//
//	return t
//}

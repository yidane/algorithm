package problems

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTree(arr []interface{}) *TreeNode {
	if len(arr) == 0 {
		return nil
	}

	if arr[0] == nil {
		return nil
	}

	root := TreeNode{Val: arr[0].(int)}
	root.appendNode(0, arr)

	return &root
}

func (root *TreeNode) appendNode(i int, arr []interface{}) {
	li := i*2 + 1
	ri := li + 1
	l := len(arr)

	if li < l {
		if arr[li] != nil {
			root.AddLeft(arr[li].(int))
			root.Left.appendNode(li, arr)
		}
	}

	if ri < l {
		if arr[ri] != nil {
			root.AddRight(arr[ri].(int))
			root.Right.appendNode(ri, arr)
		}
	}
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

func (root *TreeNode) Equal(root1 *TreeNode) bool {
	if root == nil {
		if root1 == nil {
			return true
		}
		return false
	} else {
		if root1 == nil {
			return false
		}
	}

	if root.Val != root1.Val {
		return false
	}

	return root.Left.Equal(root1.Left) && root.Right.Equal(root1.Right)
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

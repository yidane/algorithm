package tree

func Add(root *TreeNode, val int) {
	if root == nil {
		root = NewTreeNode(val)
	}

	if val > root.Val {
		if root.Right == nil {
			root.Right = NewTreeNode(val)
			return
		}

		Add(root.Right, val)
	} else {
		if root.Left == nil {
			root.Left = NewTreeNode(val)
			return
		}

		Add(root.Left, val)
	}
}

func Max(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Right == nil {
		return root.Val
	}

	return Max(root.Right)
}

func Min(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil {
		return root.Val
	}

	return Min(root.Left)
}

func Search(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == val {
		return root
	}

	if val > root.Val {
		return Search(root.Right, val)
	} else {
		return Search(root.Left, val)
	}
}

func Delete(root *TreeNode, val int) {
	if root == nil {
		return
	}

	for {
		if root.Val == val {
			if root.Left == nil {
				if root.Right == nil { //左右节点都为nil,该节点为叶子节点，直接删除
					root = nil
				} else { //右节点不为nil,右节点直接赋值给该节点
					root = root.Right
				}
				break
			} else {
				if root.Right == nil { //右节点不为nil,左节点直接赋值给该节点
					root = root.Left
					break
				} else {
					//左右节点都不为nil，怎么处理？
					//右节点上提

					root = root.Right
				}
			}
		}

		if val > root.Val {
			root = root.Right
		} else {
			root = root.Left
		}
	}
}

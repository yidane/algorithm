package problems

import "testing"

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		root func() *TreeNode
		want bool
	}{
		{root: func() *TreeNode {
			return NewTreeNode(2).AddLeft(1).AddRight(3)
		}, want: true},
		{root: func() *TreeNode {
			node := NewTreeNode(5)
			l := NewTreeNode(1).AddLeft(3).AddRight(6)
			node.Left = l
			node.AddRight(4)

			return node
		}, want: false},
		{root: func() *TreeNode {
			return NewTreeNode(1).AddLeft(1)
		}, want: false},
		{root: func() *TreeNode {
			//[10,5,15,null,null,6,20]
			node := NewTreeNode(10).AddLeft(5).AddRight(15)
			node.Right.AddLeft(6).AddRight(20)

			return node
		}, want: false},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := isValidBST(tt.root()); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

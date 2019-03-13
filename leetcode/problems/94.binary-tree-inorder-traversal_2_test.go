package problems

import (
	"reflect"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
	tests := []struct {
		name string
		args func() *TreeNode
		want []int
	}{
		{name: "", args: func() *TreeNode {
			node := NewTreeNode(1)
			node.AddRight(2)
			node.Right.AddLeft(3)

			return node

		}, want: []int{1, 3, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversal(tt.args()); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal() = %v, want %v", got, tt.want)
			}

			if got := inorderTraversal1(tt.args()); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

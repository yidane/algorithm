package problems

import (
	"testing"
)

func TestNewTree(t *testing.T) {
	arr := []interface{}{1, nil, 2}

	root := NewTree(arr)

	if root == nil {
		t.Fail()
	}

	if root.Left != nil {
		t.Fail()
	}

	if root.Right == nil {
		t.Fail()
	}
}

func TestTreeNode_Equal(t *testing.T) {
	tests := []struct {
		arr1 []interface{}
		arr2 []interface{}
		want bool
	}{
		{arr1: []interface{}{1, 2, 3, 4, 5}, arr2: []interface{}{1, 2, 3, 4, 5}, want: true},
		{arr1: []interface{}{1, 2, nil, 4, 5}, arr2: []interface{}{1, nil, 3, 4, 5}, want: false},
		{arr1: []interface{}{}, arr2: []interface{}{1}, want: false},
		{arr1: []interface{}{}, arr2: []interface{}{}, want: true},
		{arr1: []interface{}{1, 2}, arr2: []interface{}{1}, want: false},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			root1 := NewTree(tt.arr1)
			root2 := NewTree(tt.arr2)
			if got := root1.Equal(root2); got != tt.want {
				t.Errorf("TreeNode.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

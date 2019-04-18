package problems

import (
	"fmt"
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

func TestPrintTree(t *testing.T) {
	arr := []interface{}{1, 2, 3, 4, 5}
	root := NewTree(arr)

	m := make(map[int][]*TreeNode)
	findLevel(root, 0, m)

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			fmt.Print(m[i][j].Val, " ")
		}

		fmt.Println()
	}
}

func findLevel(node *TreeNode, l int, m map[int][]*TreeNode) {
	if node == nil {
		return
	}

	m[l] = append(m[l], node)
	findLevel(node.Left, l+1, m)
	findLevel(node.Right, l+1, m)
}

func TestPrintTree2(t *testing.T) {
	arr := []interface{}{1, 2, 3, 4, 5}
	root := NewTree(arr)
	nodes := []*TreeNode{root}

	for len(nodes) > 0 {
		newNodes := make([]*TreeNode, 0)
		for i := 0; i < len(nodes); i++ {
			node := nodes[i]
			fmt.Print(node.Val, " ")

			if node.Left != nil {
				newNodes = append(newNodes, node.Left)
			}

			if node.Right != nil {
				newNodes = append(newNodes, node.Right)
			}
		}
		fmt.Println()

		nodes = newNodes
	}
}

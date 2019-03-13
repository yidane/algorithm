package problems

import (
	"testing"
)

func Test_deleteNode(t *testing.T) {
	type testCase struct {
		nodes *ListNode
		val   int
		want  []int
	}

	findNode := func(node *ListNode, v int) *ListNode {
		for node != nil && node.Val != v {
			node = node.Next
		}

		if node == nil {
			panic("no such node")
		}

		return node
	}

	tests := []testCase{
		{nodes: NewListNode([]int{4, 5, 1, 9}), val: 4, want: []int{5, 1, 9}},
		{nodes: NewListNode([]int{4, 5, 1, 9}), val: 5, want: []int{4, 1, 9}},
		{nodes: NewListNode([]int{4, 5, 1, 9}), val: 1, want: []int{4, 5, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.nodes.StringSequence(), func(t *testing.T) {
			deleteNode(findNode(tt.nodes, tt.val))

			if !tt.nodes.Equal(tt.want) {
				t.Fatal(tt.nodes.ToArray(), tt.want)
			}
		})
	}
}

package problems

import (
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	tests := []struct {
		head *ListNode
		n    int
		want []int
	}{
		{head: NewListNode([]int{1, 2, 3, 4, 5}), n: 1, want: []int{1, 2, 3, 4}},
		{head: NewListNode([]int{1, 2, 3, 4, 5}), n: 2, want: []int{1, 2, 3, 5}},
		{head: NewListNode([]int{1, 2, 3, 4, 5}), n: 3, want: []int{1, 2, 4, 5}},
		{head: NewListNode([]int{1, 2, 3, 4, 5}), n: 4, want: []int{1, 3, 4, 5}},
		{head: NewListNode([]int{1, 2, 3, 4, 5}), n: 5, want: []int{2, 3, 4, 5}},
		{head: NewListNode([]int{1}), n: 1, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.head.StringSequence(), func(t *testing.T) {
			if got := removeNthFromEnd(tt.head, tt.n); !got.Equal(tt.want) {
				t.Errorf("removeNthFromEnd() = %v,%v, want %v", got.StringSequence(), tt.n, tt.want)
			}
		})
	}
}

package problems

import (
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	tests := []struct {
		args [][]int
		want []int
	}{
		{args: [][]int{
			{1, 4, 5},
			{1, 3, 4},
			{2, 6},
		}, want: []int{1, 1, 2, 3, 4, 4, 5, 6}},
		{args: [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8}}, want: []int{1, 2, 3, 4, 5, 6, 7, 8}},
		{args: [][]int{
			{1, 2, 3},
		}, want: []int{1, 2, 3}},
		{args: [][]int{
			{1}, {2}, {3},
		}, want: []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			var lists []*ListNode

			for i := 0; i < len(tt.args); i++ {
				lists = append(lists, NewListNode(tt.args[i]))
			}

			if got := mergeKLists(lists); !got.Equal(tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}

			var lists1 []*ListNode

			for i := 0; i < len(tt.args); i++ {
				lists1 = append(lists1, NewListNode(tt.args[i]))
			}

			if got := mergeKLists1(lists1); !got.Equal(tt.want) {
				t.Errorf("mergeKLists1() = %v, want %v", got, tt.want)
			}
		})
	}
}

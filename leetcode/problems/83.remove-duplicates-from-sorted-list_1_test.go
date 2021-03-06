package problems

import (
	"fmt"
	"testing"
)

func Test_deleteDuplicates1(t *testing.T) {
	tests := []struct {
		args []int
		want []int
	}{
		{args: []int{1, 2, 3}, want: []int{1, 2, 3}},
		{args: []int{1, 2, 3, 3}, want: []int{1, 2, 3}},
		{args: []int{1, 2, 2, 3}, want: []int{1, 2, 3}},
		{args: []int{1, 2, 3, 3, 4, 4, 5}, want: []int{1, 2, 3, 4, 5}},
		{args: []int{1, 1, 1, 2, 3}, want: []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			head := NewListNode(tt.args)
			if got := deleteDuplicates1(head); !got.Equal(tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", got.StringSequence(), tt.want)
			}
		})
	}
}

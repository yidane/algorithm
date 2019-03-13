package problems

import (
	"fmt"
	"testing"
)

func Test_swapPairs(t *testing.T) {
	tests := []struct {
		args []int
		want []int
	}{
		{args: []int{1, 2, 3, 4}, want: []int{2, 1, 4, 3}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			head := NewListNode(tt.args)
			if got := swapPairs(head); !got.Equal(tt.want) {
				t.Errorf("swapPairs() = %v, want %v", got.StringSequence(), tt.want)
			}
		})
	}
}

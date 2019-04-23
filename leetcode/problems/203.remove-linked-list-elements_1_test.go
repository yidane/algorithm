package problems

import (
	"fmt"
	"testing"
)

func Test_removeElements(t *testing.T) {
	tests := []struct {
		args []int
		val  int
		want []int
	}{
		{[]int{1, 2, 6, 3, 4, 5, 6}, 6, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 2, 1}, 2, []int{1, 1}},
		{[]int{1, 1, 1, 2, 3}, 1, []int{2, 3}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			head := NewListNode(tt.args)

			got := removeElements(head, tt.val)
			if !got.Equal(tt.want) {
				t.Errorf("removeElements() = %v, want %v", got.ToArray(), tt.want)
			}
		})
	}
}

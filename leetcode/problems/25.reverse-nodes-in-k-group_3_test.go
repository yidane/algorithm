package problems

import (
	"fmt"
	"testing"
)

func Test_reverseKGroup(t *testing.T) {
	tests := []struct {
		args []int
		k    int
		want []int
	}{
		{args: []int{1, 2, 3, 4, 5}, k: 2, want: []int{2, 1, 4, 3, 5}},
		{args: []int{1, 2, 3, 4, 5}, k: 3, want: []int{3, 2, 1, 4, 5}},
		{args: []int{1, 2, 3, 4, 5}, k: 4, want: []int{4, 3, 2, 1, 5}},
		{args: []int{1, 2, 3, 4, 5}, k: 5, want: []int{5, 4, 3, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			head := NewListNode(tt.args)

			if got := reverseKGroup(head, tt.k); !got.Equal(tt.want) {
				t.Errorf("reverseKGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

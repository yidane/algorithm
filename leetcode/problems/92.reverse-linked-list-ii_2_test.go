package problems

import (
	"fmt"
	"testing"
)

func Test_reverseBetween(t *testing.T) {
	tests := []struct {
		args []int
		m    int
		n    int
		want []int
	}{
		{args: []int{1, 2, 3, 4, 5}, m: 1, n: 1, want: []int{1, 2, 3, 4, 5}},
		{args: []int{1, 2, 3, 4, 5}, m: 1, n: 2, want: []int{2, 1, 3, 4, 5}},
		{args: []int{1, 2, 3, 4, 5}, m: 2, n: 3, want: []int{1, 3, 2, 4, 5}},
		{args: []int{1, 2, 3, 4, 5}, m: 2, n: 4, want: []int{1, 4, 3, 2, 5}},
		{args: []int{1, 2, 3, 4, 5}, m: 3, n: 5, want: []int{1, 2, 5, 4, 3}},
		{args: []int{1, 2, 3, 4, 5}, m: 5, n: 5, want: []int{1, 2, 3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args, tt.m, tt.n), func(t *testing.T) {
			head := NewListNode(tt.args)
			if got := reverseBetween(head, tt.m, tt.n); !got.Equal(tt.want) {
				t.Errorf("reverseBetween() = %v, want %v", got.ToArray(), tt.want)
			}
		})
	}
}

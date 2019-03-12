package problems

import (
	"fmt"
	"testing"
)

func Test_partition(t *testing.T) {
	tests := []struct {
		args []int
		x    int
		want []int
	}{
		{args: []int{1, 4, 3, 2, 5, 2}, x: 3, want: []int{1, 2, 2, 4, 3, 5}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			head := NewListNode(tt.args)
			if got := partition(head, tt.x); !got.Equal(tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}

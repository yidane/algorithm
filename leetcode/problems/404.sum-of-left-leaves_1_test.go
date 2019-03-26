package problems

import (
	"fmt"
	"testing"
)

func Test_sumOfLeftLeaves(t *testing.T) {
	tests := []struct {
		args []interface{}
		want int
	}{
		{args: []interface{}{3, 9, 20, nil, nil, 15, 7}, want: 24},
		{args: []interface{}{1, 2, 3, 4, 5}, want: 4},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			root := NewTree(tt.args)

			if got := sumOfLeftLeaves(root); got != tt.want {
				t.Errorf("sumOfLeftLeaves() = %v, want %v", got, tt.want)
			}
		})
	}
}

package problems

import (
	"fmt"
	"testing"
)

func Test_minDepth(t *testing.T) {
	tests := []struct {
		args []interface{}
		want int
	}{
		{[]interface{}{1, nil, 2}, 2},
		{[]interface{}{3, 9, 20, nil, nil, 15, 7}, 2},
		{[]interface{}{1, 2, 3, 4, nil, nil, 5}, 3},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			root := NewTree(tt.args)
			if got := minDepth(root); got != tt.want {
				t.Errorf("minDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}

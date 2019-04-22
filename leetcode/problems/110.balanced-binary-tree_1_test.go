package problems

import (
	"fmt"
	"testing"
)

func Test_isBalanced(t *testing.T) {
	tests := []struct {
		args []interface{}
		want bool
	}{
		{[]interface{}{3, 9, 20, nil, nil, 15, 7}, true},
		{[]interface{}{1, 2, 2, 3, 3, nil, nil, 4, 4}, false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			root := NewTree(tt.args)

			if got := isBalanced(root); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}

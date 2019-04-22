package problems

import (
	"fmt"
	"testing"
)

func Test_hasPathSum(t *testing.T) {
	tests := []struct {
		args []interface{}
		sum  int
		want bool
	}{
		{[]interface{}{1, 2}, 1, false},
		{[]interface{}{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, nil, 1}, 22, true},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			root := NewTree(tt.args)
			if got := hasPathSum(root, tt.sum); got != tt.want {
				t.Errorf("hasPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

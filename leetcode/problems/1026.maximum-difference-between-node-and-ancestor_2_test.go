package problems

import (
	"fmt"
	"testing"
)

func Test_maxAncestorDiff(t *testing.T) {
	tests := []struct {
		args []interface{}
		want int
	}{
		{args: []interface{}{8, 3, 10, 1, 6, nil, 14, nil, nil, 4, 7, 13}, want: 7},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			root := NewTree(tt.args)

			if got := maxAncestorDiff(root); got != tt.want {
				t.Errorf("maxAncestorDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}

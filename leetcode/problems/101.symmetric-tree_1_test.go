package problems

import (
	"fmt"
	"testing"
)

func Test_isSymmetric(t *testing.T) {
	tests := []struct {
		args []interface{}
		want bool
	}{
		{[]interface{}{1, 2, 2, 3, 4, 4, 3}, true},
		{[]interface{}{1, 2, 2, nil, 3, nil, 3}, false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			root := NewTree(tt.args)

			if got := isSymmetric(root); got != tt.want {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}

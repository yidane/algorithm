package problems

import (
	"fmt"
	"testing"
)

func Test_sortedArrayToBST(t *testing.T) {
	tests := []struct {
		args []int
	}{
		{args: []int{-10, -3, 0, 5, 9}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			root := sortedArrayToBST(tt.args)

			fmt.Println(root)
		})
	}
}

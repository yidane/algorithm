package problems

import (
	"fmt"
	"testing"
)

func Test_findTilt(t *testing.T) {
	tests := []struct {
		args []interface{}
		want int
	}{
		{args: []interface{}{1, 2, 3}, want: 1},
		{args: []interface{}{1, 2, 3, 4, nil, 5}, want: 11},
		{args: []interface{}{1, 2, nil, 3, nil, 4, nil, 5}, want: 23},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			root := NewTree(tt.args)

			if got := findTilt(root); got != tt.want {
				t.Errorf("findTilt() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*

    	1
	2		3
4		 5
*/

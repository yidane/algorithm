package problems

import (
	"fmt"
	"testing"
)

func Test_rob(t *testing.T) {
	tests := []struct {
		args []int
		want int
	}{
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{2, 1, 1, 2}, 4},
		{[]int{1, 2, 3, 1}, 4},
		{[]int{2, 7, 9, 3, 1}, 12},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			if got := rob(tt.args); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}

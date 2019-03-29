package problems

import (
	"fmt"
	"testing"
)

func Test_thirdMax(t *testing.T) {
	tests := []struct {
		args []int
		want int
	}{
		{[]int{3, 2, 1}, 1},
		{[]int{1, 2}, 2},
		{[]int{1, 1, 2}, 2},
		{[]int{2, 2, 3, 1}, 1},
		{[]int{3, 3, 4, 3, 0, 0, 4, 3}, 0},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			if got := thirdMax(tt.args); got != tt.want {
				t.Errorf("thirdMax() = %v, want %v", got, tt.want)
			}
		})
	}
}

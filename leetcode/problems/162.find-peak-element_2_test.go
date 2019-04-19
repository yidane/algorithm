package problems

import (
	"fmt"
	"testing"
)

func Test_findPeakElement(t *testing.T) {
	tests := []struct {
		args []int
		want int
	}{
		{[]int{1}, 0},
		{[]int{1, 2}, 1},
		{[]int{1, 2, 1}, 1},
		{[]int{1, 2, 3}, 2},
		{[]int{1, 2, 3, 2}, 2},
		{[]int{1, 2, 3, 4}, 3},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			if got := findPeakElement(tt.args); got != tt.want {
				t.Errorf("findPeakElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

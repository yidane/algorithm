package problems

import (
	"fmt"
	"testing"
)

func Test_search(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 5, 1},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{[]int{1, 3, 5}, 3, 1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.nums, tt.target), func(t *testing.T) {
			if got := search33(tt.nums, tt.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}

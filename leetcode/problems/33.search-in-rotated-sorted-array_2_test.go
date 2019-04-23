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
		{nil, 1, -1},
		{[]int{}, 1, -1},
		{[]int{1}, 1, 0},
		{[]int{0, 1, 2, 4, 5, 6, 7}, 0, 0},
		{[]int{0, 1, 2, 4, 5, 6, 7}, 3, -1},
		{[]int{7, 0, 1, 2, 4, 5, 6}, 0, 1},
		{[]int{7, 0, 1, 2, 4, 5, 6}, 3, -1},
		{[]int{6, 7, 0, 1, 2, 4, 5}, 0, 2},
		{[]int{6, 7, 0, 1, 2, 4, 5}, 3, -1},
		{[]int{5, 6, 7, 0, 1, 2, 4}, 0, 3},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{2, 4, 5, 6, 7, 0, 1}, 3, -1},
		{[]int{2, 4, 5, 6, 7, 0, 1}, 0, 5},
		{[]int{1, 2, 4, 5, 6, 7, 0}, 3, -1},
		{[]int{1, 2, 4, 5, 6, 7, 0}, 0, 6},
		{[]int{5, 6, 7, 0, 1, 2, 4}, 3, -1},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 5, 1},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{[]int{4, 5, 6, 7, 8, 1, 2, 3}, 8, 4},
		{[]int{4, 5, 6, 7, 8, 9, 0, 1, 2, 3}, 0, 6},
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

package problems

import (
	"fmt"
	"testing"
)

func Test_maxProfit121(t *testing.T) {
	tests := []struct {
		prices []int
		want   int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 5, 4, 3, 2, 1}, 0},
		{[]int{7, 1, 8, 2, 9, 3}, 8},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.prices), func(t *testing.T) {
			if got := maxProfit121(tt.prices); got != tt.want {
				t.Errorf("maxProfit121() = %v, want %v", got, tt.want)
			}
		})
	}
}

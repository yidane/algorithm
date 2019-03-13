package problems

import (
	"fmt"
	"testing"
)

func Test_subarraySum(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		//{nums: []int{1, 1, 1}, k: 2, want: 2},
		{nums: []int{1, 1, 1, 1}, k: 2, want: 3},
		{nums: []int{-1, 1, 0, -1, 1, 0}, k: 0, want: 11},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.nums), func(t *testing.T) {
			if got := subarraySum(tt.nums, tt.k); got != tt.want {
				t.Errorf("subarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}

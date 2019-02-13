package problems

import (
	"fmt"
	"testing"
)

func Test_rotate(t *testing.T) {
	type args struct {
		nums []int
		k    int
		want []int
	}
	tests := []args{
		{nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 3, want: []int{5, 6, 7, 1, 2, 3, 4}},
		{nums: []int{-1, -100, 3, 99}, k: 2, want: []int{3, 99, -1, -100}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			rotate1(tt.nums, tt.k)

			fmt.Println(tt.nums)
		})
	}
}

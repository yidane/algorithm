package problems

import (
	"fmt"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums  []int
		wants []int
	}
	tests := []args{
		{nums: []int{0, 1, 0, 3, 12}, wants: []int{1, 3, 12, 0, 0}},
		{nums: []int{0, 0, 1}, wants: []int{1, 0, 0}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			moveZeroes1(tt.nums)

			fmt.Println(tt.nums)
			for i := 0; i < len(tt.nums); i++ {
				if tt.nums[i] != tt.wants[i] {
					t.Fatalf("%v should be equal %v", tt.nums[i], tt.wants[i])
				}
			}
		})
	}
}

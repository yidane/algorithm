package problems

import (
	"fmt"
	"testing"
)

func Test_merge(t *testing.T) {
	tests := []struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}{
		{[]int{0}, 0, []int{1}, 1, []int{1}},
		{[]int{2, 0}, 1, []int{1}, 1, []int{1, 2}},
		{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
		{[]int{1, 2, 4, 5, 6, 0}, 5, []int{3}, 1, []int{1, 2, 3, 4, 5, 6}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.nums1), func(t *testing.T) {
			merge(tt.nums1, tt.m, tt.nums2, tt.n)

			if !SameArray(tt.nums1, tt.want) {
				t.Fatal(tt.nums1, tt.want)
			}
		})
	}
}

package problems

import (
	"fmt"
	"testing"
)

func Test_findContentChildren(t *testing.T) {
	tests := []struct {
		g    []int
		s    []int
		want int
	}{
		//{[]int{1, 2, 3}, []int{1, 1}, 1},
		{[]int{1, 2}, []int{1, 2, 3}, 2},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.g, tt.s), func(t *testing.T) {
			if got := findContentChildren(tt.g, tt.s); got != tt.want {
				t.Errorf("findContentChildren() = %v, want %v", got, tt.want)
			}
		})
	}
}

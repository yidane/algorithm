package problems

import (
	"fmt"
	"testing"
)

func Test_largestPerimeter(t *testing.T) {
	tests := []struct {
		args []int
		want int
	}{
		{args: []int{2, 1, 2}, want: 5},
		{args: []int{1, 2, 1}, want: 0},
		{args: []int{3, 2, 3, 4}, want: 10},
		{args: []int{3, 6, 2, 3}, want: 8},
		{args: []int{1, 2, 2, 4, 18, 8}, want: 5},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			if got := largestPerimeter(tt.args); got != tt.want {
				t.Errorf("largestPerimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}

package problems

import (
	"strconv"
	"testing"
)

func Test_majorityElement(t *testing.T) {
	tests := []struct {
		args []int
		want int
	}{
		//{args: []int{3, 2, 3}, want: 3},
		{args: []int{2, 2, 1, 1, 1, 2, 2}, want: 2},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.want), func(t *testing.T) {
			if got := majorityElement(tt.args); got != tt.want {
				t.Errorf("majorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

package problems

import (
	"strconv"
	"testing"
)

func Test_climbStairs(t *testing.T) {
	tests := []struct {
		arg  int
		want int
	}{
		{arg: 1, want: 1},
		{arg: 2, want: 2},
		{arg: 3, want: 3},
		{arg: 4, want: 5},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.arg), func(t *testing.T) {
			if got := climbStairs(tt.arg); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

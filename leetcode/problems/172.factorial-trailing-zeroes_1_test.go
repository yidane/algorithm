package problems

import (
	"fmt"
	"testing"
)

func Test_trailingZeroes(t *testing.T) {
	tests := []struct {
		args int
		want int
	}{
		{1, 0},
		{2, 0},
		{3, 0},
		{4, 0},
		{5, 1},
		{6, 1},
		{7, 1},
		{8, 1},
		{9, 1},
		{10, 2},
		{30, 7},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			if got := trailingZeroes(tt.args); got != tt.want {
				t.Errorf("trailingZeroes() = %v, want %v", got, tt.want)
			}

			if got := trailingZeroes1(tt.args); got != tt.want {
				t.Errorf("trailingZeroes() = %v, want %v", got, tt.want)
			}
		})
	}
}

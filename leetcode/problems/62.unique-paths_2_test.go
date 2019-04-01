package problems

import (
	"fmt"
	"testing"
)

func Test_uniquePaths(t *testing.T) {
	tests := []struct {
		m    int
		n    int
		want int
	}{
		{3, 2, 3},
		{7, 3, 28},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.m, tt.n, tt.want), func(t *testing.T) {
			if got := uniquePaths(tt.m, tt.n); got != tt.want {
				t.Errorf("uniquePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}

package problems

import (
	"fmt"
	"testing"
)

func Test_matrixScore(t *testing.T) {
	tests := []struct {
		args [][]int
		want int
	}{
		{args: [][]int{{0, 0, 1, 1}, {1, 0, 1, 0}, {1, 1, 0, 0}}, want: 39},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			if got := matrixScore(tt.args); got != tt.want {
				t.Errorf("matrixScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

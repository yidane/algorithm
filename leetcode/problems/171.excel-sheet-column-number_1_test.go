package problems

import (
	"testing"
)

func Test_titleToNumber(t *testing.T) {
	tests := []struct {
		args string
		want int
	}{
		{"A", 1},
		{"B", 2},
		{"Z", 26},
		{"AA", 27},
		{"AB", 28},
		{"ZY", 701},
	}
	for _, tt := range tests {
		t.Run(tt.args, func(t *testing.T) {
			if got := titleToNumber(tt.args); got != tt.want {
				t.Errorf("titleToNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

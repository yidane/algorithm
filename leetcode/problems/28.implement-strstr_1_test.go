package problems

import (
	"fmt"
	"testing"
)

func Test_strStr(t *testing.T) {
	tests := []struct {
		haystack string
		needle   string
		want     int
	}{
		{"hello", "ll", 2},
		{"hello", "lll", -1},
		{"hello", "", 0},
		{"hello", "o", 4},
		{"hello", "lo", 3},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.haystack, "IndexOf", tt.needle), func(t *testing.T) {
			if got := strStr(tt.haystack, tt.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

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

func Test_KMP(t *testing.T) {
	tests := []struct {
		haystack string
		needle   string
		want     int
	}{
		{"BBC ABCDAB ABCDABDE", "ABCDABD", -1},
		//{"hello", "ll", 2},
		//{"hello", "lll", -1},
		//{"hello", "", 0},
		//{"hello", "o", 4},
		//{"hello", "lo", 3},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.haystack, "IndexOf", tt.needle), func(t *testing.T) {
			if got := KMP(tt.haystack, tt.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_kmpNext(t *testing.T) {
	tests := []struct {
		args string
		want []int
	}{
		{args: "ABCDABD", want: []int{0, 0, 0, 0, 1, 2, 0}},
		{args: "ABCDABDD", want: []int{0, 0, 0, 0, 1, 2, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.args, func(t *testing.T) {
			if got := kmpNext(tt.args); !SameArray(got, tt.want) {
				t.Errorf("kmpNext() = %v, want %v", got, tt.want)
			}
		})
	}
}

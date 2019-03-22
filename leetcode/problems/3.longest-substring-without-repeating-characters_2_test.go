package problems

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		args string
		want int
	}{
		{"abcabcabc", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"abba", 2},
		{"uqinntq", 4},
		{"abcdefghihabckefgmn", 11},
	}
	for _, tt := range tests {
		t.Run(tt.args, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

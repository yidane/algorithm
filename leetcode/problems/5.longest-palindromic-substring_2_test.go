package problems

import "testing"

func Test_longestPalindrome(t *testing.T) {
	tests := []struct {
		args string
		want int
	}{
		{"", 0},
		{"a", 1},
		{"babad", 3},
		{"aaaaaa", 6},
		{"babad", 3},
		{"cbbd", 2},
		{"abcba", 5},
		{"abcda", 1},
		{"aaabaaaa", 7},
		{"aacdefcaa", 2},
	}
	for _, tt := range tests {
		t.Run(tt.args, func(t *testing.T) {
			got := longestPalindrome(tt.args)
			if len(got) != tt.want && isPalindromeString(got) {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

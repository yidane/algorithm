package problems

import "testing"

func Test_validPalindrome(t *testing.T) {
	tests := []struct {
		name string
		args string
		want bool
	}{
		{name: "aba", args: "aba", want: true},
		{name: "abca", args: "abca", want: true},
		{name: "abab", args: "abab", want: true},
		{name: "ababc", args: "ababc", want: false},
		{name: "abadba", args: "abadba", want: true},
		{name: "abadcffcdba", args: "abadcffcdba", want: true},
		{name: "ebcbbececabbacecbbcbe", args: "ebcbbececabbacecbbcbe", want: true},
		{name: "aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga", args: "aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga", want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validPalindrome(tt.args); got != tt.want {
				t.Errorf("validPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

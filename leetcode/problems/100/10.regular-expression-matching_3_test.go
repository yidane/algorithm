package problems

import "testing"

func Test_isMatch(t *testing.T) {
	tests := []struct {
		s    string
		p    string
		want bool
	}{
		{s: "aa", p: "a", want: false},
		{s: "aa", p: "a*", want: true},
		{s: "ab", p: ".*", want: true},
		{s: "aab", p: "c*a*b", want: true},
		{s: "mississippi", p: "mis*is*p*.", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := isMatch(tt.s, tt.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

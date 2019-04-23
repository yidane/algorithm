package problems

import (
	"fmt"
	"testing"
)

func Test_isIsomorphic(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want bool
	}{
		{"egg", "add", true},
		{"foo", "bar", false},
		{"paper", "title", true},
		{"ab", "aa", false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.s, tt.t), func(t *testing.T) {
			if got := isIsomorphic(tt.s, tt.t); got != tt.want {
				t.Errorf("isIsomorphic() = %v, want %v", got, tt.want)
			}
		})
	}
}

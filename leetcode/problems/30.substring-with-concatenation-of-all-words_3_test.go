package problems

import (
	"testing"
)

func Test_findSubstring(t *testing.T) {
	tests := []struct {
		s     string
		words []string
		want  []int
	}{
		//{"barfoothefoobarman", []string{"foo", "bar"}, []int{0, 9}},
		//{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}, []int{}},
		//{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}, []int{8}},
		{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := findSubstring(tt.s, tt.words); !SameArray(got, tt.want) {
				t.Errorf("findSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

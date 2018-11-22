package problems

import "testing"

func TestMagicDictionary_Search(t *testing.T) {
	tests := []struct {
		name string
		word string
		want bool
	}{
		{name: "hello", word: "hello", want: false},
		{name: "hhllo", word: "hhllo", want: true},
		{name: "hell", word: "hell", want: false},
		{name: "leetcoded", word: "leetcoded", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := Constructor()
			this.BuildDict([]string{"hello", "leetcode"})
			if got := this.Search(tt.word); got != tt.want {
				t.Errorf("%v MagicDictionary.Search() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
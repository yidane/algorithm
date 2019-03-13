package problems

import "testing"

func Test_lengthOfLastWord(t *testing.T) {
	tests := map[string]int{
		"":             0,
		"yidane":       6,
		"yidane ":      6,
		"hello world":  5,
		"hello world ": 5,
	}
	for k, v := range tests {
		t.Run(k, func(t *testing.T) {
			if got := lengthOfLastWord(k); got != v {
				t.Errorf("lengthOfLastWord(%s) = %v, want %v", k, got, v)
			}
		})
	}
}

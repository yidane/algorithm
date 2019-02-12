package problems

import "testing"

func Test_isValid(t *testing.T) {
	tests := map[string]bool{
		"()":     true,
		"(":      false,
		"1":      false,
		"()[]{}": true,
		"([)]":   false,
		"{[]}":   true,
		"(]":     false,
		")[":     false,
	}
	for k, v := range tests {
		t.Run(k, func(t *testing.T) {
			if got := isValid(k); got != v {
				t.Errorf("isValid(%s) = %v, want %v", k, got, v)
			}
		})
	}
}

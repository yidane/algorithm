package problems

import "testing"

func Test_isPalindrome(t *testing.T) {
	tests := map[int]bool{
		1:    true,
		-1:   false,
		10:   false,
		11:   true,
		12:   false,
		13:   false,
		21:   false,
		100:  false,
		111:  true,
		1001: true,
	}
	for k, v := range tests {
		t.Run("", func(t *testing.T) {
			if got := isPalindrome(k); got != v {
				t.Errorf("isPalindrome(%v) = %v, want %v", k, got, v)
			}
		})
	}
}

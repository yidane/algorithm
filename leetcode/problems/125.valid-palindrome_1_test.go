package problems

import (
	"testing"
)

func Test_isPalindrome125(t *testing.T) {
	tests := []struct {
		args string
		want bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
	}
	for _, tt := range tests {
		t.Run(tt.args, func(t *testing.T) {
			if got := isPalindrome125(tt.args); got != tt.want {
				t.Errorf("isPalindrome125() = %v, want %v", got, tt.want)
			}

			if got := isPalindrome125_1(tt.args); got != tt.want {
				t.Errorf("isPalindrome125() = %v, want %v", got, tt.want)
			}
		})
	}
}

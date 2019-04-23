package problems

import (
	"fmt"
	"testing"
)

func Test_isPowerOfTwo(t *testing.T) {
	tests := []struct {
		args int
		want bool
	}{
		{1, true},
		{16, true},
		{218, false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			if got := isPowerOfTwo(tt.args); got != tt.want {
				t.Errorf("isPowerOfTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}

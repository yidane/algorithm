package problems

import (
	"strconv"
	"testing"
)

func Test_addDigits(t *testing.T) {
	tests := []struct {
		args int
		want int
	}{
		{args: 38, want: 2},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.args), func(t *testing.T) {
			if got := addDigits(tt.args); got != tt.want {
				t.Errorf("addDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}

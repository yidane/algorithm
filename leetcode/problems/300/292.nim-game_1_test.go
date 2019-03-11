package problems

import (
	"strconv"
	"testing"
)

func Test_canWinNim(t *testing.T) {
	tests := []struct {
		args int
		want bool
	}{
		{args: 1, want: true},
		{args: 2, want: true},
		{args: 3, want: true},
		{args: 4, want: false},
		{args: 5, want: true},
		{args: 6, want: true},
		{args: 7, want: true},
		{args: 8, want: false},
		{args: 9, want: true},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.args), func(t *testing.T) {
			if got := canWinNim(tt.args); got != tt.want {
				t.Errorf("canWinNim() = %v, want %v", got, tt.want)
			}
		})
	}
}

package problems

import (
	"fmt"
	"testing"
)

func Test_divide(t *testing.T) {
	tests := []struct {
		dividend int
		divisor  int
		want     int
	}{
		{dividend: 10, divisor: 3, want: 3},
		{dividend: 7, divisor: -3, want: -2},
		{dividend: -1, divisor: -1, want: 1},
		{dividend: -2147483648, divisor: -1, want: 2147483647},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.dividend, tt.divisor), func(t *testing.T) {
			if got := divide(tt.dividend, tt.divisor); got != tt.want {
				t.Errorf("divide() = %v, want %v", got, tt.want)
			}
		})
	}
}

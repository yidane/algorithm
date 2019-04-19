package problems

import (
	"fmt"
	"testing"
)

func Test_fractionToDecimal(t *testing.T) {
	tests := []struct {
		numerator   int
		denominator int
		want        string
	}{
		{1, 2, "0.5"},
		{2, 1, "2"},
		{2, 3, "0.(6)"},
		{2, 3, "0.(6)"},
		{10, 7, "1.(428571)"},
		{-50, 8, "-6.25"},
		{7, -12, "-0.58(3)"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.numerator, tt.denominator), func(t *testing.T) {
			if got := fractionToDecimal(tt.numerator, tt.denominator); got != tt.want {
				t.Errorf("fractionToDecimal() = %v, want %v", got, tt.want)
			}
		})
	}
}

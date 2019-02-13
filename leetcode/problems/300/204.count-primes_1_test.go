package problems

import (
	"strconv"
	"testing"
)

func Test_countPrimes(t *testing.T) {
	type args struct {
		n    int
		want int
	}
	tests := []args{
		{n: 2, want: 0},
		{n: 10, want: 4},
		{n: 12, want: 5},
		{n: 13, want: 5},
		{n: 999983, want: 5},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.n), func(t *testing.T) {
			if got := countPrimes(tt.n); got != tt.want {
				t.Errorf("countPrimes() = %v, want %v", got, tt.want)
			}
		})
	}
}

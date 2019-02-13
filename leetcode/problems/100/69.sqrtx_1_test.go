package problems

import "testing"

func Test_mySqrt(t *testing.T) {
	type args struct {
		x    int
		want int
	}
	tests := []args{
		{x: 2, want: 1},
		{x: 3, want: 1},
		{x: 4, want: 2},
		{x: 8, want: 2},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := mySqrt(tt.x); got != tt.want {
				t.Errorf("mySqrt(%v) = %v, want %v", tt.x, got, tt.want)
			}
		})
	}
}

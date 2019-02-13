package problems

import "testing"

func Test_isPowerOfFour(t *testing.T) {
	type args struct {
		num  int
		want bool
	}
	tests := []args{
		{num: 16, want: true},
		{num: 17, want: false},
		{num: 1, want: true},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := isPowerOfFour(tt.num); got != tt.want {
				t.Errorf("isPowerOfFour() = %v, want %v", got, tt.want)
			}
		})
	}
}

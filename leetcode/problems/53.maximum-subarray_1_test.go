package problems

import "testing"

func Test_maxSubArray(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		//{args: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, want: 6},
		//{args: []int{2, 1, -1, 5, 1, -5, 4}, want: 8},
		//{args: []int{-2, 1, -1, 5, 1, -5, 4}, want: 6},
		{args: []int{-2, 3, -1, 5, 1, -5, 4}, want: 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.args); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

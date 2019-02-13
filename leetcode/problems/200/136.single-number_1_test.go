package problems

import "testing"

func Test_singleNumber(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{args: []int{2, 2, 1}, want: 1},
		{args: []int{4, 1, 2, 1, 2}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber(tt.args); got != tt.want {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

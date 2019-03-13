package problems

import "testing"

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		args []int
		want int
	}{
		//{args: []int{}, want: 0},
		{args: []int{1, 1, 2}, want: 2},
		{args: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, want: 5},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := removeDuplicates(tt.args); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

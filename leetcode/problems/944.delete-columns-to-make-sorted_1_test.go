package problems

import (
	"fmt"
	"testing"
)

func Test_minDeletionSize(t *testing.T) {
	tests := []struct {
		args []string
		want int
	}{
		{args: []string{"zyx", "wvu", "tsr"}, want: 3},
		{args: []string{"a", "b"}, want: 0},
		{args: []string{"cba", "daf", "ghi"}, want: 1},
	}
	for _, tt := range tests {
		args1 := make([]string, len(tt.args))
		copy(args1, tt.args)
		t.Run(fmt.Sprint(args1), func(t *testing.T) {
			if got := minDeletionSize(args1); got != tt.want {
				t.Errorf("minDeletionSize() = %v, want %v", got, tt.want)
			}
		})

		args2 := make([]string, len(tt.args))
		copy(args2, tt.args)
		t.Run(fmt.Sprint(args2), func(t *testing.T) {
			if got := minDeletionSize1(args2); got != tt.want {
				t.Errorf("minDeletionSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

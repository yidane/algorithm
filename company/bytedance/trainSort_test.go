package bytedance

import (
	"fmt"
	"testing"
)

func Test_trainSort(t *testing.T) {
	tests := []struct {
		in   int
		out  []int
		want bool
	}{
		{2, []int{2, 1}, true},
		{5, []int{3, 4, 1, 5, 2}, true},
		{5, []int{3, 4, 2, 1, 5}, false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.out), func(t *testing.T) {
			if got := trainSort(tt.in, tt.out); got != tt.want {
				t.Errorf("trainSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

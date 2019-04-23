package problems

import (
	"fmt"
	"testing"
)

func Test_containsNearbyDuplicate(t *testing.T) {
	tests := []struct {
		args []int
		k    int
		want bool
	}{
		{[]int{1, 2, 3, 1}, 3, true},
		{[]int{1, 0, 1, 1}, 1, true},
		{[]int{1, 2, 3, 1, 2, 3}, 2, false},
		{[]int{4, 1, 2, 3, 1, 5}, 3, true},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			if got := containsNearbyDuplicate(tt.args, tt.k); got != tt.want {
				t.Errorf("containsNearbyDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}

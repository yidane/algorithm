package problems

import (
	"fmt"
	"testing"
)

func Test_largestSumAfterKNegations(t *testing.T) {
	tests := []struct {
		args []int
		k    int
		want int
	}{
		//{args: []int{4, 2, 3}, k: 1, want: 5},
		//{args: []int{3, -1, 0, 2}, k: 3, want: 6},
		//{args: []int{2, -3, -1, 5, -4}, k: 2, want: 13},
		{args: []int{-2, 5, 0, 2, -2}, k: 3, want: 11},
	}
	for _, tt := range tests {
		args1 := make([]int, len(tt.args))
		copy(args1, tt.args)
		t.Run(fmt.Sprint(args1), func(t *testing.T) {
			if got := largestSumAfterKNegations(args1, tt.k); got != tt.want {
				t.Errorf("largestSumAfterKNegations() = %v, want %v", got, tt.want)
			}
		})

		args2 := make([]int, len(tt.args))
		copy(args2, tt.args)
		t.Run(fmt.Sprint(args2), func(t *testing.T) {
			if got := largestSumAfterKNegations1(args2, tt.k); got != tt.want {
				t.Errorf("largestSumAfterKNegations() = %v, want %v", got, tt.want)
			}
		})

		args3 := make([]int, len(tt.args))
		copy(args3, tt.args)
		t.Run(fmt.Sprint(args3), func(t *testing.T) {
			if got := largestSumAfterKNegations2(args3, tt.k); got != tt.want {
				t.Errorf("largestSumAfterKNegations() = %v, want %v", got, tt.want)
			}
		})
	}
}

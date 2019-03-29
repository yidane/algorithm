package problems

import (
	"fmt"
	"testing"
)

func Test_findDisappearedNumbers(t *testing.T) {
	tests := []struct {
		args []int
		want []int
	}{
		{args: []int{4, 3, 2, 7, 8, 2, 3, 1}, want: []int{5, 6}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			if got := findDisappearedNumbers(tt.args); !SameArray(got, tt.want) {
				t.Errorf("findDisappearedNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findDisappearedNumbers1(t *testing.T) {
	tests := []struct {
		args []int
		want []int
	}{
		{args: []int{4, 3, 2, 7, 8, 2, 3, 1}, want: []int{5, 6}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			if got := findDisappearedNumbers1(tt.args); !SameArray(got, tt.want) {
				t.Errorf("findDisappearedNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

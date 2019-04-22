package problems

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_plusOne(t *testing.T) {
	tests := []struct {
		args []int
		want []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
		{[]int{8}, []int{9}},
		{[]int{9}, []int{1, 0}},
		{[]int{9, 9}, []int{1, 0, 0}},
		{[]int{9, 9, 9}, []int{1, 0, 0, 0}},
		{[]int{8, 9, 9, 9}, []int{9, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			if got := plusOne(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

package problems

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_twoSum167(t *testing.T) {
	tests := []struct {
		numbers []int
		target  int
		want    []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.numbers), func(t *testing.T) {
			if got := twoSum167(tt.numbers, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

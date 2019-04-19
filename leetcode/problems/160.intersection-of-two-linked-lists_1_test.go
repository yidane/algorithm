package problems

import (
	"fmt"
	"testing"
)

func Test_getIntersectionNode(t *testing.T) {
	tests := []struct {
		arr1 []int
		arr2 []int
		want []int
	}{
		{[]int{2, 6, 4}, []int{1, 5}, nil},
		{[]int{0, 9, 1, 2, 4}, []int{3, 2, 4}, []int{2, 4}},
		{[]int{0, 9, 1, 2, 1, 4}, []int{3, 2, 4}, []int{4}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.want), func(t *testing.T) {
			headA := NewListNode(tt.arr1)
			headB := NewListNode(tt.arr2)

			got := getIntersectionNode(headA, headB)

			if !got.Equal(tt.want) {
				t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

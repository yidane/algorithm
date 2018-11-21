package problems

import (
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	tests := []struct {
		arr1 []int
		arr2 []int
		arr3 []int
	}{
		{arr1: []int{5}, arr2: []int{5}, arr3: []int{0, 1}},
	}

	for _, test := range tests {
		l1 := NewListNode(test.arr1)
		l2 := NewListNode(test.arr2)
		l3 := addTwoNumbers(l1, l2)

		if !l3.Equal(test.arr3) {
			t.Fatal(test)
		}
	}
}

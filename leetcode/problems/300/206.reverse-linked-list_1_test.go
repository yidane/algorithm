package problems

import (
	"fmt"
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head []int
		want []int
	}
	tests := []args{
		{head: []int{1, 2, 3, 4, 5}, want: []int{5, 4, 3, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.head), func(t *testing.T) {
			head := NewListNode(tt.head)

			got := reverseList(head)
			if !got.Equal(tt.want) {
				t.Fatalf("result of %v should be equal %v", tt.head, got.Val)
			}
		})

		t.Run(fmt.Sprint(tt.head), func(t *testing.T) {
			head := NewListNode(tt.head)

			got := reverseList1(head)
			if !got.Equal(tt.want) {
				t.Fatalf("result of %v should be equal %v", tt.head, got.Val)
			}
		})
	}
}

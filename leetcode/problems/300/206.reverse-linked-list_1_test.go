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
			var head *ListNode
			if len(tt.head) > 0 {
				head = new(ListNode)
				head.Val = tt.head[0]

				next := head
				for i := 1; i < len(tt.head); i++ {
					newNode := new(ListNode)
					newNode.Val = tt.head[i]
					next.Next = newNode
					next = newNode
				}
			}

			got := reverseList(head)

			if len(tt.want) == 0 {
				if got != nil {
					t.Fatalf("result of %v should be nil", tt.head)
				}
			} else {
				for i := 0; i < len(tt.want); i++ {
					if got.Val != tt.want[i] {
						t.Fatalf("result of %v should be equal %v but %v", tt.head, tt.want[i], got.Val)
					}
				}
			}
		})
	}
}

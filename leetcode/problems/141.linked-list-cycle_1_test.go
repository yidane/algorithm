package problems

import (
	"fmt"
	"testing"
)

func Test_hasCycle(t *testing.T) {
	tests := []struct {
		args []int
		pos  int
		want bool
	}{
		{[]int{3, 2, 0, 4}, 1, true},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			head := NewListNode(tt.args)

			tail := head
			for tail.Next != nil {
				tail = tail.Next
			}

			if tt.pos != -1 {
				posNode := head
				for tt.pos >= 0 {
					if posNode.Next != nil {
						posNode = posNode.Next
					}
					tt.pos--
				}

				tail.Next = posNode
			}

			if got := hasCycle(head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

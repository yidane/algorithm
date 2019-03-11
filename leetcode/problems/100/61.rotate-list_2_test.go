package problems

import (
	"testing"
)

func Test_rotateRight(t *testing.T) {

	type testCase struct {
		heads []int
		k     int
		wants []int
	}

	testCases := []testCase{
		{heads: []int{1, 2, 3, 4, 5}, k: 2, wants: []int{4, 5, 1, 2, 3}},
		{heads: []int{}, k: 2, wants: []int{}},
		{heads: []int{1}, k: 2, wants: []int{1}},
		{heads: []int{1, 3}, k: 2, wants: []int{1, 3}},
		{heads: []int{0, 1, 2}, k: 4, wants: []int{2, 0, 1}},
		{heads: []int{0, 1}, k: 0, wants: []int{0, 1}},
		{heads: []int{1, 2}, k: 2, wants: []int{1, 2}},
	}

	for _, tt := range testCases {
		t.Run("", func(t *testing.T) {
			//prepare ListNode
			var head *ListNode
			if len(tt.heads) != 0 {
				head = new(ListNode)
				head.Val = tt.heads[0]

				n := head
				for i := 1; i < len(tt.heads); i++ {
					n.Next = new(ListNode)
					n.Next.Val = tt.heads[i]
					n = n.Next
				}
			}

			//rotate ListNode
			r := rotateRight(head, tt.k)

			//assert
			if len(tt.wants) == 0 {
				if r != nil {
					t.Fatal("the result should not be nil")
				}
			} else {
				for i := 0; i < len(tt.wants); i++ {
					if r == nil {
						t.Fatal("the result should not be nil")
					}

					if r.Val != tt.wants[i] {
						t.Fatalf("%v should equal %v", r.Val, tt.wants[i])
					}

					r = r.Next
				}
			}
		})
	}
}

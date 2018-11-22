package problems

import (
	"fmt"
	"github.com/yidane/algorithm/leetcode/problems/data"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	l1 := data.ListNode{Val: 2}
	l2 := data.ListNode{Val: 1}
	fmt.Println(mergeTwoLists(&l1, &l2))
}

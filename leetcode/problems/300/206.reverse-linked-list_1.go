package problems

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	n := head
	var node *ListNode
	for n.Next != nil {
		fmt.Println(n.Val)
		t := n
		t.Next = node
		node = t
		n = n.Next
	}

	head.Next = nil
	return node
}

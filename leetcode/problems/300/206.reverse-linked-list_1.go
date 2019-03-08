package problems

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 1 2 3 4 5
// 2 1 3 4 5
// 3 2 1 4 5
// 4 3 2 1 5
// 5 4 3 2 1

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	h := head
	n := h.Next
	for n != nil {
		t := n.Next
		n.Next = h
		head.Next = t
		h = n

		if t == nil {
			return h
		}

		n = t
	}

	return h
}

func reverseList1(head *ListNode) *ListNode {
	var prevNode *ListNode
	prevNode = nil
	currentNode := head
	for currentNode != nil {
		nextRef := currentNode.Next
		currentNode.Next = prevNode
		prevNode = currentNode
		currentNode = nextRef
	}
	return prevNode
}

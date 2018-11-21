package problems

/*
You are given two non-empty linked lists representing two non-negative integers.
The digits are stored in reverse order and each of their nodes contain a single digit.
Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	newNode := new(ListNode)
	rtnNode := newNode
	var nextPlus = false
	newNode.Val = l1.Val + l2.Val
	if newNode.Val > 9 {
		newNode.Val = newNode.Val % 10
		nextPlus = true
	}

	l1 = l1.Next
	l2 = l2.Next

	for l1 != nil || l2 != nil {
		newNode.Next = new(ListNode)

		switch {
		case l1 == nil:
			newNode.Next.Val = l2.Val
			l2 = l2.Next
		case l2 == nil:
			newNode.Next.Val = l1.Val
			l1 = l1.Next
		default:
			newNode.Next.Val = l1.Val + l2.Val
			l1 = l1.Next
			l2 = l2.Next
		}

		if nextPlus {
			newNode.Next.Val++
			nextPlus = false
		}

		if newNode.Next.Val > 9 {
			newNode.Next.Val = newNode.Next.Val % 10
			nextPlus = true
		}

		newNode = newNode.Next
	}

	if nextPlus {
		newNode.Next = new(ListNode)
		newNode.Next.Val = 1
	}

	return rtnNode
}

package problems

import "github.com/yidane/algorithm/leetcode/problems/data"

//*************************************************************************
/*
https://leetcode.com/problems/merge-two-sorted-lists/#/description

Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.
*/
//*************************************************************************

func mergeTwoLists(l1 *data.ListNode, l2 *data.ListNode) *data.ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	result := new(data.ListNode)
	next := result
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			next.Next = l1
			l1 = l1.Next
		} else {
			next.Next = l2
			l2 = l2.Next
		}
		next = next.Next
	}
	if l1 != nil {
		next.Next = l1
	} else {
		next.Next = l2
	}

	return result.Next
}

/*
 	var head *ListNode = new(ListNode)
	var next *ListNode = head
	for nil != l1 && nil != l2 {
		if l1.Val  < l2.Val {
			next.Next = l1
			l1 = l1.Next
		}else{
			next.Next = l2
			l2 = l2.Next
		}
		next = next.Next
	}
	if nil != l1{
		next.Next = l1
	}else{
		next.Next = l2
	}
	return head.Next
*/

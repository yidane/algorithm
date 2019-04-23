package problems

/*
 * @lc app=leetcode id=203 lang=golang
 *
 * [203] Remove Linked List Elements
 *
 * https://leetcode.com/problems/remove-linked-list-elements/description/
 *
 * algorithms
 * Easy (35.57%)
 * Total Accepted:    218.1K
 * Total Submissions: 612.5K
 * Testcase Example:  '[1,2,6,3,4,5,6]\n6'
 *
 * Remove all elements from a linked list of integers that have value val.
 *
 * Example:
 *
 *
 * Input:  1->2->6->3->4->5->6, val = 6
 * Output: 1->2->3->4->5
 *
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	for head != nil {
		if head.Val != val {
			break
		}

		head = head.Next
	}

	if head == nil {
		return nil
	}

	node := head
	for node != nil && node.Next != nil {
		if node.Next.Val == val {
			node.Next = node.Next.Next
			continue
		}
		node = node.Next
	}

	return head
}

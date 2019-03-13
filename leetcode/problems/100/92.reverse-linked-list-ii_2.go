package problems

/*
 * @lc app=leetcode id=92 lang=golang
 *
 * [92] Reverse Linked List II
 *
 * https://leetcode.com/problems/reverse-linked-list-ii/description/
 *
 * algorithms
 * Medium (34.14%)
 * Total Accepted:    180.9K
 * Total Submissions: 529.8K
 * Testcase Example:  '[1,2,3,4,5]\n2\n4'
 *
 * Reverse a linked list from position m to n. Do it in one-pass.
 *
 * Note: 1 ≤ m ≤ n ≤ length of list.
 *
 * Example:
 *
 *
 * Input: 1->2->3->4->5->NULL, m = 2, n = 4
 * Output: 1->4->3->2->5->NULL
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
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n || head == nil {
		return head
	}

	i := 1

	h := head
	e := h
	head = head.Next
	h.Next = nil
	for head != nil {
		cur := head
		head = head.Next
		cur.Next = nil

		i++
		if i < m {
			e.Next = cur
			continue
		}

		if i <= n {
			cur.Next = h
			h = cur
			continue
		}

		e.Next = cur
		e = e.Next
	}

	return h
}

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

1 2 3 4 5
1 3 2 4 5
1 4 3 2 5

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

	var prevHead = new(ListNode)
	var result = prevHead
	var h, e *ListNode
	for i := 1; head != nil; i++ {
		if i < m {
			prevHead.Next = head
			prevHead = prevHead.Next
			head = head.Next
			continue
		}

		if i > n {
			break
		}

		cur := head
		head = head.Next

		if h == nil {
			cur.Next = nil
			h = cur
			e = cur
		} else {
			cur.Next = h
			h = cur
		}
	}

	prevHead.Next = h
	e.Next = head

	return result.Next
}

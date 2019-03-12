package problems

/*
 * @lc app=leetcode id=25 lang=golang
 *
 * [25] Reverse Nodes in k-Group
 *
 * https://leetcode.com/problems/reverse-nodes-in-k-group/description/
 *
 * algorithms
 * Hard (35.45%)
 * Total Accepted:    170.1K
 * Total Submissions: 479.8K
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * Given a linked list, reverse the nodes of a linked list k at a time and
 * return its modified list.
 *
 * k is a positive integer and is less than or equal to the length of the
 * linked list. If the number of nodes is not a multiple of k then left-out
 * nodes in the end should remain as it is.
 *
 *
 *
 *
 * Example:
 *
 * Given this linked list: 1->2->3->4->5
 *
 * For k = 2, you should return: 2->1->4->3->5
 *
 * For k = 3, you should return: 3->2->1->4->5
 *
 * Note:
 *
 *
 * Only constant extra memory is allowed.
 * You may not alter the values in the list's nodes, only nodes itself may be
 * changed.
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
func reverseKGroup(head *ListNode, k int) *ListNode {
	var result = new(ListNode)
	var resultHead = result
	for head != nil {
		var begin, end *ListNode = head, nil
		i := 0
		for ; i < k; i++ {
			if head == nil {
				break
			}

			head = head.Next
			end = head
		}

		if i == k {
			var prev *ListNode
			var groupTail = begin
			for {
				if begin == end {
					break
				}

				next := begin.Next
				begin.Next = prev
				prev = begin
				begin = next
			}

			result.Next = prev
			result = groupTail
		} else {
			result.Next = begin
			break
		}
	}

	return resultHead.Next
}

package problems

/*
 * @lc app=leetcode id=24 lang=golang
 *
 * [24] Swap Nodes in Pairs
 *
 * https://leetcode.com/problems/swap-nodes-in-pairs/description/
 *
 * algorithms
 * Medium (43.33%)
 * Total Accepted:    285.5K
 * Total Submissions: 658.8K
 * Testcase Example:  '[1,2,3,4]'
 *
 * Given a linked list, swap every two adjacent nodes and return its head.
 *
 * You may not modify the values in the list's nodes, only nodes itself may be
 * changed.
 *
 *
 *
 * Example:
 *
 *
 * Given 1->2->3->4, you should return the list as 2->1->4->3.
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
func swapPairs(head *ListNode) *ListNode {
	m := make(map[int]*ListNode)
	i := 1
	for {
		if head == nil {
			break
		}

		m[i] = head
		head = head.Next
		i++
	}

	var result = new(ListNode)
	head = result
	for j := 1; j < i; j++ {
		if j+1 < i {
			head.Next = m[j+1]
			head = head.Next
		}

		head.Next = m[j]
		head = head.Next
		j++
	}

	head.Next = nil

	return result.Next
}

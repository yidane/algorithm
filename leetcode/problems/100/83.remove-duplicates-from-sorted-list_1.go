package problems

/*
 * @lc app=leetcode id=83 lang=golang
 *
 * [83] Remove Duplicates from Sorted List
 *
 * https://leetcode.com/problems/remove-duplicates-from-sorted-list/description/
 *
 * algorithms
 * Easy (41.96%)
 * Total Accepted:    304.7K
 * Total Submissions: 726.2K
 * Testcase Example:  '[1,1,2]'
 *
 * Given a sorted linked list, delete all duplicates such that each element
 * appear only once.
 *
 * Example 1:
 *
 *
 * Input: 1->1->2
 * Output: 1->2
 *
 *
 * Example 2:
 *
 *
 * Input: 1->1->2->3->3
 * Output: 1->2->3
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

//sorted linklist ...
func deleteDuplicates1(head *ListNode) *ListNode {
	h1 := head
	m := make(map[int]bool)

	for h1 != nil {
		if _, ok := m[h1.Val]; !ok {
			m[h1.Val] = true
		}

		h1 = h1.Next
	}

	var result = new(ListNode)
	var resultHead = result
	for head != nil {
		cur := head
		head = head.Next
		cur.Next = nil
		if f := m[cur.Val]; f {
			m[cur.Val] = false
			result.Next = cur
			result = result.Next
		}
	}

	return resultHead.Next
}

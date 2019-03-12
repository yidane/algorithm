package problems

/*
 * @lc app=leetcode id=82 lang=golang
 *
 * [82] Remove Duplicates from Sorted List II
 *
 * https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/description/
 *
 * algorithms
 * Medium (32.26%)
 * Total Accepted:    170.1K
 * Total Submissions: 527.1K
 * Testcase Example:  '[1,2,3,3,4,4,5]'
 *
 * Given a sorted linked list, delete all nodes that have duplicate numbers,
 * leaving only distinct numbers from the original list.
 *
 * Example 1:
 *
 *
 * Input: 1->2->3->3->4->4->5
 * Output: 1->2->5
 *
 *
 * Example 2:
 *
 *
 * Input: 1->1->1->2->3
 * Output: 2->3
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
func deleteDuplicates(head *ListNode) *ListNode {
	h1 := head
	m := make(map[int]int)

	containDuplicate := false
	for h1 != nil {
		if _, ok := m[h1.Val]; ok {
			m[h1.Val] = m[h1.Val] + 1
			containDuplicate = true
		} else {
			m[h1.Val] = 1
		}

		h1 = h1.Next
	}

	if !containDuplicate {
		return head
	}

	var result = new(ListNode)
	var resultHead = result
	for head != nil {
		cur := head
		head = head.Next
		cur.Next = nil
		if v := m[cur.Val]; v == 1 {

			result.Next = cur
			result = result.Next
		}
	}

	return resultHead.Next
}

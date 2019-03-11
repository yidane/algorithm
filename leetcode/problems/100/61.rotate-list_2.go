package problems

/*
 * @lc app=leetcode id=61 lang=golang
 *
 * [61] Rotate List
 *
 * https://leetcode.com/problems/rotate-list/description/
 *
 * algorithms
 * Medium (26.56%)
 * Total Accepted:    179.4K
 * Total Submissions: 675.4K
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * Given a linked list, rotate the list to the right by k places, where k is
 * non-negative.
 *
 * Example 1:
 *
 *
 * Input: 1->2->3->4->5->NULL, k = 2
 * Output: 4->5->1->2->3->NULL
 * Explanation:
 * rotate 1 steps to the right: 5->1->2->3->4->NULL
 * rotate 2 steps to the right: 4->5->1->2->3->NULL
 *
 *
 * Example 2:
 *
 *
 * Input: 0->1->2->NULL, k = 4
 * Output: 2->0->1->NULL
 * Explanation:
 * rotate 1 steps to the right: 2->0->1->NULL
 * rotate 2 steps to the right: 1->2->0->NULL
 * rotate 3 steps to the right: 0->1->2->NULL
 * rotate 4 steps to the right: 2->0->1->NULL
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	if k <= 0 {
		return head
	}

	end := head

	sum := 1
	for end.Next != nil {
		end = end.Next
		sum++
	}

	k = k % sum
	if k == 0 {
		return head
	}

	//set as ring
	end.Next = head

	for sum-k > 1 {
		head = head.Next
		k++
	}
	r := head.Next
	head.Next = nil

	return r
}

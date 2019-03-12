package problems

/*
 * @lc app=leetcode id=86 lang=golang
 *
 * [86] Partition List
 *
 * https://leetcode.com/problems/partition-list/description/
 *
 * algorithms
 * Medium (36.42%)
 * Total Accepted:    154.2K
 * Total Submissions: 423.5K
 * Testcase Example:  '[1,4,3,2,5,2]\n3'
 *
 * Given a linked list and a value x, partition it such that all nodes less
 * than x come before nodes greater than or equal to x.
 *
 * You should preserve the original relative order of the nodes in each of the
 * two partitions.
 *
 * Example:
 *
 *
 * Input: head = 1->4->3->2->5->2, x = 3
 * Output: 1->2->2->4->3->5
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
func partition(head *ListNode, x int) *ListNode {
	leftHead := new(ListNode)
	leftTail := leftHead
	rightHead := new(ListNode)
	rightTail := rightHead

	for head != nil {
		cur := head
		head = head.Next
		cur.Next = nil

		if cur.Val >= x {
			rightTail.Next = cur
			rightTail = rightTail.Next
		} else {
			leftTail.Next = cur
			leftTail = leftTail.Next
		}
	}

	leftTail.Next = rightHead.Next

	return leftHead.Next
}

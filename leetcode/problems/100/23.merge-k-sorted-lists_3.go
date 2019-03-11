package problems

import (
	"sort"
)

/*
 * @lc app=leetcode id=23 lang=golang
 *
 * [23] Merge k Sorted Lists
 *
 * https://leetcode.com/problems/merge-k-sorted-lists/description/
 *
 * algorithms
 * Hard (33.14%)
 * Total Accepted:    347.5K
 * Total Submissions: 1M
 * Testcase Example:  '[[1,4,5],[1,3,4],[2,6]]'
 *
 * Merge k sorted linked lists and return it as one sorted list. Analyze and
 * describe its complexity.
 *
 * Example:
 *
 *
 * Input:
 * [
 * 1->4->5,
 * 1->3->4,
 * 2->6
 * ]
 * Output: 1->1->2->3->4->4->5->6
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

//节约内存
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 1 {
		return lists[0]
	}

	head := new(ListNode)
	var result = head
	var cur *ListNode
	complete := false
	t := 0
	for {
		complete = true
		for i := 0; i < len(lists); i++ {
			if lists[i] == nil {
				continue
			}

			complete = false

			if cur == nil {
				cur = lists[i]
				t = i
				continue
			}

			if lists[i].Val < cur.Val {
				cur = lists[i]
				t = i
			}
		}

		head.Next = cur
		head = head.Next

		if complete {
			break
		}

		cur = cur.Next

		if lists[t] != nil {
			lists[t] = lists[t].Next
		}
	}

	return result.Next
}

func mergeKLists1(lists []*ListNode) *ListNode {
	nums := make([]int, 0)

	for i := 0; i < len(lists); i++ {
		node := lists[i]
		for node != nil {
			nums = append(nums, node.Val)
			node = node.Next
		}
	}

	sort.Ints(nums)

	var result = new(ListNode)
	var head = result
	for i := 0; i < len(nums); i++ {
		head.Next = &ListNode{
			Val: nums[i],
		}

		head = head.Next
	}

	return result.Next
}

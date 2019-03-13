package problems

import (
	"container/list"
	"sort"
)

/*
 * @lc app=leetcode id=703 lang=golang
 *
 * [703] Kth Largest Element in a Stream
 *
 * https://leetcode.com/problems/kth-largest-element-in-a-stream/description/
 *
 * algorithms
 * Easy (45.26%)
 * Total Accepted:    24.6K
 * Total Submissions: 54.3K
 * Testcase Example:  '["KthLargest","add","add","add","add","add"]\n[[3,[4,5,8,2]],[3],[5],[10],[9],[4]]'
 *
 * Design a class to find the kth largest element in a stream. Note that it is
 * the kth largest element in the sorted order, not the kth distinct element.
 *
 * Your KthLargest class will have a constructor which accepts an integer k and
 * an integer array nums, which contains initial elements from the stream. For
 * each call to the method KthLargest.add, return the element representing the
 * kth largest element in the stream.
 *
 * Example:
 *
 *
 * int k = 3;
 * int[] arr = [4,5,8,2];
 * KthLargest kthLargest = new KthLargest(3, arr);
 * kthLargest.add(3);   // returns 4
 * kthLargest.add(5);   // returns 5
 * kthLargest.add(10);  // returns 5
 * kthLargest.add(9);   // returns 8
 * kthLargest.add(4);   // returns 8
 *
 *
 * Note:
 * You may assume that nums' length ≥ k-1 and k ≥ 1.
 *
 */
type KthLargest struct {
	l *list.List
	k int
}

func NewKthLargest(k int, nums []int) KthLargest {
	kl := KthLargest{k: k, l: list.New()}

	if len(nums) != 0 {
		sort.Ints(nums)
		for i := len(nums) - 1; i >= len(nums)-k && i >= 0; i-- {
			kl.l.PushBack(nums[i])
		}
	}

	return kl
}

func (this *KthLargest) Add(val int) int {
	if this.l.Len() == 0 {
		this.l.PushBack(val)

		return val
	}

	h := this.l.Front()
	e := this.l.Back()

	if e.Value.(int) >= val {
		this.l.PushBack(val)
	} else if h.Value.(int) <= val {
		this.l.PushFront(val)
	} else {
		for {
			if h == nil {
				this.l.PushBack(val)
				break
			}

			v := h.Value.(int)
			if v <= val {
				this.l.InsertBefore(val, h)
				break
			}

			h = h.Next()
		}
	}

	if this.l.Len() > this.k {
		this.l.Remove(this.l.Back())
	}

	return this.l.Back().Value.(int)
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

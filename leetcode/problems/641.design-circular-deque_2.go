package problems

import (
	"sync/atomic"
)

/*
 * @lc app=leetcode id=641 lang=golang
 *
 * [641] Design Circular Deque
 *
 * https://leetcode.com/problems/design-circular-deque/description/
 *
 * algorithms
 * Medium (49.11%)
 * Total Accepted:    5.9K
 * Total Submissions: 12K
 * Testcase Example:  '["MyCircularDeque","insertLast","insertLast","insertFront","insertFront","getRear","isFull","deleteLast","insertFront","getFront"]\n[[3],[1],[2],[3],[4],[],[],[],[4],[]]'
 *
 * Design your implementation of the circular double-ended queue (deque).
 *
 * Your implementation should support following operations:
 *
 *
 * MyCircularDeque(k): Constructor, set the size of the deque to be k.
 * insertFront(): Adds an item at the front of Deque. Return true if the
 * operation is successful.
 * insertLast(): Adds an item at the rear of Deque. Return true if the
 * operation is successful.
 * deleteFront(): Deletes an item from the front of Deque. Return true if the
 * operation is successful.
 * deleteLast(): Deletes an item from the rear of Deque. Return true if the
 * operation is successful.
 * getFront(): Gets the front item from the Deque. If the deque is empty,
 * return -1.
 * getRear(): Gets the last item from Deque. If the deque is empty, return
 * -1.
 * isEmpty(): Checks whether Deque is empty or not.
 * isFull(): Checks whether Deque is full or not.
 *
 *
 *
 *
 * Example:
 *
 *
 * MyCircularDeque circularDeque = new MycircularDeque(3); // set the size to
 * be 3
 * circularDeque.insertLast(1);            // return true
 * circularDeque.insertLast(2);            // return true
 * circularDeque.insertFront(3);            // return true
 * circularDeque.insertFront(4);            // return false, the queue is full
 * circularDeque.getRear();              // return 2
 * circularDeque.isFull();                // return true
 * circularDeque.deleteLast();            // return true
 * circularDeque.insertFront(4);            // return true
 * circularDeque.getFront();            // return 4
 *
 *
 *
 *
 * Note:
 *
 *
 * All values will be in the range of [0, 1000].
 * The number of operations will be in the range of [1, 1000].
 * Please do not use the built-in Deque library.
 *
 *
 */

type dequeItem struct {
	prev *dequeItem
	next *dequeItem
	v    int
}

type MyCircularDeque struct {
	c    uint64
	l    uint64
	head *dequeItem
	tail *dequeItem
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func MyCircularDequeConstructor(k int) MyCircularDeque {
	if k < 0 {
		k = 0
	}

	return MyCircularDeque{
		c: uint64(k),
		l: 0,
	}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (m *MyCircularDeque) InsertFront(value int) bool {
	if m.l >= m.c {
		return false
	}

	newItem := &dequeItem{
		v: value,
	}

	if m.head == nil {
		m.head = newItem
		m.tail = newItem
	} else {
		newItem.next = m.head
		m.head = newItem
	}
	m.l += 1

	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (m *MyCircularDeque) InsertLast(value int) bool {
	if m.l == m.c {
		return false
	}

	newItem := &dequeItem{
		v: value,
	}

	if m.tail == nil {
		m.head = newItem
		m.tail = newItem
	} else {
		m.tail.next = newItem
		m.tail = newItem
	}

	m.l += 1

	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (m *MyCircularDeque) DeleteFront() bool {
	if atomic.LoadUint64(&m.l) == 0 {
		return false
	}

	m.head = m.head.next
	if m.head == nil {
		m.tail = nil
	}
	m.l -= 1

	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (m *MyCircularDeque) DeleteLast() bool {
	if atomic.LoadUint64(&m.l) == 0 {
		return false
	}

	if m.tail == nil {
		return true
	}

	if m.tail.prev == nil {
		m.head = nil
		m.tail = nil
	} else {
		m.tail = m.tail.prev
	}

	m.l -= 1

	return true
}

/** Get the front item from the deque. */
func (m *MyCircularDeque) GetFront() int {
	if m.head == nil {
		return -1
	}
	return m.head.v
}

/** Get the last item from the deque. */
func (m *MyCircularDeque) GetRear() int {
	if m.tail == nil {
		return -1
	}
	return m.tail.v
}

/** Checks whether the circular deque is empty or not. */
func (m *MyCircularDeque) IsEmpty() bool {
	return atomic.LoadUint64(&m.l) == 0
}

/** Checks whether the circular deque is full or not. */
func (m *MyCircularDeque) IsFull() bool {
	return atomic.LoadUint64(&m.l) == atomic.LoadUint64(&m.c)
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */

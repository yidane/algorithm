/*
 * @lc app=leetcode id=146 lang=golang
 *
 * [146] LRU Cache
 *
 * https://leetcode.com/problems/lru-cache/description/
 *
 * algorithms
 * Hard (23.93%)
 * Total Accepted:    260.9K
 * Total Submissions: 1.1M
 * Testcase Example:  '["LRUCache","put","put","get","put","get","put","get","get","get"]\n[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]'
 *
 *
 * Design and implement a data structure for Least Recently Used (LRU) cache.
 * It should support the following operations: get and put.
 *
 *
 *
 * get(key) - Get the value (will always be positive) of the key if the key
 * exists in the cache, otherwise return -1.
 * put(key, value) - Set or insert the value if the key is not already present.
 * When the cache reached its capacity, it should invalidate the least recently
 * used item before inserting a new item.
 *
 *
 * Follow up:
 * Could you do both operations in O(1) time complexity?
 */

// Example:
// LRUCache cache = new LRUCache( 2 /* capacity */ );
/*
 * cache.put(1, 1);
 * cache.put(2, 2);
 * cache.get(1);       // returns 1
 * cache.put(3, 3);    // evicts key 2
 * cache.get(2);       // returns -1 (not found)
 * cache.put(4, 4);    // evicts key 1
 * cache.get(1);       // returns -1 (not found)
 * cache.get(3);       // returns 3
 * cache.get(4);       // returns 4
 */

package problems

type item struct {
	k    int
	v    int
	next *item
	prev *item
}

type LRUCache struct {
	cap  int
	size int
	p    map[int]*item
	head *item
	tail *item
}

func Constructor(capacity int) LRUCache {
	return LRUCache{cap: capacity, p: make(map[int]*item, capacity)}
}

func (this *LRUCache) Get(key int) int {
	if oldItem, ok := this.p[key]; ok {
		this.active(oldItem)
		return oldItem.v
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	currentItem, ok := this.p[key]
	if ok {
		currentItem.v = value
	} else {
		currentItem = &item{k: key, v: value}

		if this.size == this.cap {
			delete(this.p, this.tail.k)
			if this.tail.prev != nil {
				this.tail.prev.next = nil
			}
			this.tail = this.tail.prev //抛弃最久没使用的
		} else {
			this.size++
		}
		this.p[key] = currentItem
	}

	this.active(currentItem)
}

func (this *LRUCache) active(e *item) {
	if this.head == nil || this.cap == 1 {
		this.tail = e
		this.head = e
		return
	}

	if e.k == this.head.k {
		return
	}

	if this.cap == 2 {
		this.tail = this.head
		this.head = e

		return
	}

	if e.k == this.tail.k { //如果是尾部元素
		this.tail = this.tail.prev
		e.next = this.head
		this.head.prev = e
		this.head = e
		return
	}

	if e.prev == nil { //新元素
		if this.head.k == this.tail.k {
			this.tail.prev = e
			e.next = this.tail
			this.head = e
			return
		}

		e.next = this.head
		this.head.prev = e
		this.head = e
	} else {
		e.prev.next = e.next
		e.next.prev = e.prev

		e.next = this.head
		this.head.prev = e
		e.prev = nil
		this.head = e
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

package problems

import (
	"bytes"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(array []int) *ListNode {
	if array == nil {
		return nil
	}

	newListNode := new(ListNode)
	rtnListNode := newListNode
	for i := 0; i < len(array); i++ {
		newListNode.Next = new(ListNode)
		newListNode.Next.Val = array[i]
		newListNode = newListNode.Next
	}

	return rtnListNode.Next
}

func (listNode *ListNode) ToArray() []int {
	length := 0
	node := listNode
	for node != nil {
		length++
		node = node.Next
	}

	arrayNode := listNode //dot not change listNode
	array := make([]int, length)
	for i := 0; i < length; i++ {
		array[i] = arrayNode.Val
		arrayNode = arrayNode.Next
	}

	return array
}

func (listNode *ListNode) Length() int {
	len1 := 0
	listNode1 := listNode
	for listNode1 != nil {
		len1++
		listNode1 = listNode1.Next
	}

	return len1
}

func (listNode *ListNode) SameAs(other *ListNode) bool {
	listNode1 := listNode
	other1 := other

	for listNode1 != nil && other1 != nil {
		if listNode1.Val != other1.Val {
			return false
		}

		listNode1 = listNode1.Next
		other1 = other1.Next
	}

	if listNode1 == nil && other1 == nil { //all must be nil at last
		return true
	}

	return false
}

func (listNode *ListNode) Equal(array []int) bool {
	if listNode == nil && array == nil {
		return true
	}

	listNode1 := listNode

	for i := 0; i < len(array); i++ {
		if listNode1 == nil {
			return false
		}

		if array[i] != listNode1.Val {
			return false
		}

		listNode1 = listNode1.Next
	}

	if listNode1 != nil {
		return false
	}

	return true
}

func (listNode *ListNode) StringSequence() string {
	buf := bytes.Buffer{}
	node := listNode

	for node != nil {
		buf.WriteString(strconv.Itoa(node.Val))
		node = node.Next
	}

	return buf.String()
}

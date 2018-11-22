package data

import (
	"fmt"
	"testing"
)

func TestListNode_ToArray(t *testing.T) {
	array := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10}
	listNode := NewListNode(array)

	array2 := listNode.ToArray()

	fmt.Println(array)
	fmt.Println(array2)

	if !SameArray(array, array2) {
		t.Error("SameArray")
	}

	if !listNode.Equal(array) {
		t.Error("Equal")
	}

	array1 := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10, 11}

	listNode1 := NewListNode(array1)

	if listNode.SameAs(listNode1) {
		t.Error("SameAs")
	}

	if listNode.Length() > listNode1.Length() {
		t.Error("Length")
	}
}

package eightSort

import (
	"testing"
)

func Test_insertionSort(t *testing.T) {
	testSort("插入排序", insertionSort, t)
}

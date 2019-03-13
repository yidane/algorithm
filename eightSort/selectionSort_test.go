package eightSort

import (
	"testing"
)

func Test_selectionSort(t *testing.T) {
	testSort("选择排序", selectionSort, t)
}

func Test_selectionSort1(t *testing.T) {
	selectionSort1([]int{3, 2, 1})
	testSort("选择排序", selectionSort1, t)
}

func Test_selectionSort2(t *testing.T) {
	selectionSort2([]int{3, 2, 1})
	testSort("选择排序", selectionSort1, t)
}

package eightSort

import (
	"testing"
)

func Test_selectionSort(t *testing.T) {
	testSort("选择排序", selectionSort, t)
}

func Test_selectionSort1(t *testing.T) {
	testSort("选择排序", selectionSort1, t)
}

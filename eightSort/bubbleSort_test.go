package eightSort

import (
	"testing"
)

func Test_bubbleSort(t *testing.T) {
	testSort("冒泡排序", bubbleSort, t)
}

func Test_bubbleSort1(t *testing.T) {
	testSort("冒泡排序1", bubbleSort1, t)
}

func Test_bubbleSort2(t *testing.T) {
	testSort("冒泡排序2", bubbleSort2, t)
}

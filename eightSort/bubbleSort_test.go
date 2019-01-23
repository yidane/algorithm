package eightSort

import (
	"testing"
)

func Test_bubbleSort(t *testing.T) {

}

func Test_bubbleSortError(t *testing.T) {
	testSort("冒泡排序", bubbleSort, t)
}

func Test_bubbleSort1(t *testing.T) {
	testSort("冒泡排序1", bubbleSort, t)
}

func Test_bubbleSort2(t *testing.T) {
	testSort("冒泡排序2", bubbleSort, t)
}

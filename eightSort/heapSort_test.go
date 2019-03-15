package eightSort

import (
	"fmt"
	"testing"
)

func Test_heapSort(t *testing.T) {
	//testSort("堆排序", heapSort, t)

	arr := []int{1, 3, 2}

	heapSort(arr)

	fmt.Println(arr)
}

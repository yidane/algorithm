package eightSort

import (
	"testing"
)

func Test_quickSort(t *testing.T) {
	//testSort("快速排序", quickSort, t)
	quickSort([]int{1, 2, 3})
	arr := quickSort([]int{2, 3, 4, 5, 1, 3})

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			t.Fail()
		}
	}
}

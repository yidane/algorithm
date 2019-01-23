package eightSort

import (
	"testing"
)

func Test_shellSort(t *testing.T) {
	testSort("希尔排序", shellSort, t)
}

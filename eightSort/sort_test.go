package eightSort

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSort(t *testing.T) {

	checkSort := func(actaul interface{}, expected ...interface{}) string {
		arr := actaul.([]int)
		name := expected[0].(string)
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				return fmt.Sprintf("%s计算后的结果不正确%v", name, arr)
			}
		}
		return ""
	}

	sortFuncs := map[string]func([]int) []int{
		"冒泡排序":  bubbleSort,
		"冒泡排序1": bubbleSort1,
		"冒泡排序2": bubbleSort2,
		"冒泡排序3": bubbleSortError,
		"插入排序":  insertionSort,
		"归并排序":  mergeSort,
		"快速排序":  quickSort,
		"基数排序":  radixSort,
		"希尔排序":  shellSort,
		"选择排序":  selectionSort,
		"选择排序1": selectionSort1,
	}

	Convey("测试排序算法", t, func() {
		for k := range sortFuncs {
			Convey(k, func() {
				fun := sortFuncs[k]
				for _, arr := range preSortedArr {
					//因为在执行完算法之后，数组就已经完全排序。后续排序算法就看不出效果了。
					//所有每次调用参数，都采用原数组的副本
					arr1 := make([]int, len(arr))
					copy(arr1, arr)
					So(fun(arr1), checkSort, k)
				}
			})
		}
	})
}

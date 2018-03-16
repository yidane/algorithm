package eightSort

/*
插入排序
	是一种最简单直观的排序算法，它的工作原理是通过构建有序序列，对于未排序数据，在已排序序列中从后向前扫描，找到相应位置并插入。
算法步骤：
	1）将第一待排序序列第一个元素看做一个有序序列，把第二个元素到最后一个元素当成是未排序序列。
	2）从头到尾依次扫描未排序序列，将扫描到的每个元素插入有序序列的适当位置。（如果待插入的元素与有序序列中的某个元素相等，则将待插入元素插入到相等元素的后面。）
*/
func insertionSort(arr []int) []int {
	var j, temp int
	for i := 0; i < len(arr); i++ {
		temp = arr[i]
		for j = i; j > 0 && arr[j-1] > temp; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = temp
	}
	return arr
}

package eightSort

/*
归并排序
	（Merge sort）是建立在归并操作上的一种有效的排序算法。该算法是采用分治法（Divide and Conquer）的一个非常典型的应用。
算法步骤：
	1. 申请空间，使其大小为两个已经排序序列之和，该空间用来存放合并后的序列
	2. 设定两个指针，最初位置分别为两个已经排序序列的起始位置
	3. 比较两个指针所指向的元素，选择相对小的元素放入到合并空间，并移动指针到下一位置
	4. 重复步骤3直到某一指针达到序列尾
	5. 将另一序列剩下的所有元素直接复制到合并序列尾
*/
func mergeSort(arr []int) []int {
	return split(arr, 0, len(arr)-1)
}

func split(arr []int, left, right int) []int {
	if left == right {
		return []int{arr[left]}
	}

	mid := (right + left) / 2

	arr1 := split(arr, left, mid)
	arr2 := split(arr, mid+1, right)

	return merge(arr1, arr2)
}

func merge(arr1, arr2 []int) []int {
	arr := make([]int, len(arr1)+len(arr2))
	i, j, m := 0, 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr[i] > arr[j] {
			arr[m] = arr[j]
			j++
		} else {
			arr[m] = arr[i]
			i++
		}
		m++
	}

	for ; i < len(arr1); i++ {
		arr[m] = arr1[i]
		m++
	}

	for ; j < len(arr2); j++ {
		arr[m] = arr2[j]
		m++
	}

	return arr
}

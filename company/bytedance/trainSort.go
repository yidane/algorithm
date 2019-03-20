package bytedance

/*
结题思路
只要队列能按照顺序依次排列成两队，那么两队肯定能接起来

以下解法，如果要考虑到大量数据处理，可以使用list优化，避免内存的分配
也可以通过判断第一个元素计算出所需要的内存
*/
func trainSort(in int, out []int) bool {
	var arrA, arrB []int

	cur := 1
	for i := 0; i < in; i++ {
		next := out[i]

		if !appendArr(&arrA, next) {
			if !appendArr(&arrB, next) {
				return false
			}
		}

		cur++
	}

	return true
}

//判断当前值能否接在数组后面，判断方式是此数组最后元素加1等于当前元素
func appendArr(arr *[]int, next int) bool {
	if len(*arr) == 0 {
		*arr = append(*arr, next)
		return true
	}

	if (*arr)[len(*arr)-1]+1 == next {
		*arr = append(*arr, next)
		return true
	}

	return false
}

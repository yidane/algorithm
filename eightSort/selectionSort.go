package eightSort

/*
选择排序(Selection sort)也是一种简单直观的排序算法。
算法步骤：
	1）首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置
	2）再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。
	3）重复第二步，直到所有元素均排序完毕。
*/

//时间复杂度 n!
func selectionSort(arr []int) []int {

	for i := 0; i < len(arr)-1; i++ {
		t := i
		for j := i + 1; j < len(arr); j++ {
			if arr[t] > arr[j] {
				t = j
			}
		}
		if t != i {
			arr[i], arr[t] = arr[t], arr[i]
		}
	}

	return arr
}

//时间复杂度 n!/2
//一次循环，找到最大和最小值，分别和两端交换，两端一起检索
func selectionSort1(arr []int) []int {
	for left, right := 0, len(arr)-1; left <= right; {
		min, max := left, right
		for i := left; i <= right; i++ {
			if arr[min] > arr[i] {
				min = i
			}

			if arr[max] < arr[i] {
				max = i
			}
		}

		arr[min], arr[left] = arr[left], arr[min]
		if left == max { //如果最大值位于最左边，由于上一步已经将最左边移位到min，因此此时最大值为min
			max = min
		}
		arr[max], arr[right] = arr[right], arr[max]

		left++
		right--
	}

	return arr
}

//时间复杂度 n!/2
//一次循环，找到最大和最小值，分别和两端交换，从左开始检索，只需要检索一半即可
func selectionSort2(arr []int) []int {
	l := len(arr)
	for i := 0; i < l/2+1; i++ {
		tMin := i
		tMax := l - 1 - i

		if tMax <= tMin {
			break
		}

		for j := i; j < l-i; j++ {
			if arr[tMin] > arr[j] {
				tMin = j
			}
			if arr[tMax] < arr[j] {
				tMax = j
			}
		}

		arr[i], arr[tMin] = arr[tMin], arr[i]
		if i == tMax { //如果最大值位于最左边，由于上一步已经将最左边移位到min，因此此时最大值为min
			tMax = tMin
		}
		arr[l-1-i], arr[tMax] = arr[tMax], arr[l-1-i]

	}

	return arr
}

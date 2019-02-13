package problems

/*
	1 2 3 4
	4 1 2 3
	3 4 1 2
	2 3 4 1
*/

//k*n
func rotate(nums []int, k int) {
	if len(nums) == 0 {
		return
	}

	k = k % len(nums)
	l := len(nums)
	for i := 0; i < k; i++ {
		for j := l - 1; j > 0; j-- {
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}
}

//
func rotate1(nums []int, k int) {
	if len(nums) == 0 {
		return
	}

	k = k % len(nums)
	l := len(nums)

	for i, j := 0, l-1; i < j; {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}

	for i, j := 0, k-1; i < j; {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}

	for i, j := k, l-1; i < j; {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

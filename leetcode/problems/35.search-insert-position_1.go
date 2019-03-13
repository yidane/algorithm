package problems

func searchInsert(nums []int, target int) int {
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			return i
		} else {
			index++
		}
	}

	return index
}

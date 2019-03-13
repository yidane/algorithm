package problems

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	subIndex := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			subIndex++
		} else {
			nums[i-subIndex+1] = nums[i+1]
		}
	}

	nums = nums[:len(nums)-subIndex]

	return len(nums)
}

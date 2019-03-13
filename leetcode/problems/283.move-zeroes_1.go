package problems

/*
Input: [0,1,0,3,12]
0 1 0 3 12
1 0 0 3 12
1 0 3 0 12
1 0 3 12 0
1 3 0 12 0
1 3 12 0 0
Output: [1,3,12,0,0]
*/
func moveZeroes(nums []int) {
	c := 0
	for i := 0; i < len(nums); i++ {
		if i+c == len(nums) {
			break
		}

		if nums[i] == 0 {
			for j := i; j < len(nums)-c-1; j++ {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
			c++

			if nums[i] == 0 {
				i--
			}
		}
	}
}

func moveZeroes1(nums []int) {
	nonZeros := 0
	for curr := 0; curr < len(nums); curr++ {
		if nums[curr] != 0 {
			nums[nonZeros], nums[curr] = nums[curr], nums[nonZeros]
			nonZeros++
		}
	}
}

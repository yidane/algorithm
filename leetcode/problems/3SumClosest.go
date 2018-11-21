package problems

import (
	"sort"
)

// for j := 0; j < i; j++ {
// 	for k := len(nums) - 1; k > i; k-- {
// 		total := nums[i] + nums[j] + nums[k]
// 		s := total - target
// 		switch {
// 		case s < 0:
// 			s = -s
// 		case s == 0:
// 			return total
// 		}
// 		if s < sub {
// 			t = total
// 			sub = s
// 		}
// 	}
// }

func threeSumClosest(array []int, target int) int {
	sub := 1 << 32
	t := 0
	sort.Ints(array)
	for k, v := range array {
		i := 0
		j := len(array) - 1
		for i < k && j > k {
			total := array[i] + array[j] + v
			switch {
			case total < target:
				i++
			case total == target:
				return total
			case total > target:
				j--
			}
			s := total - target
			if s < 0 {
				s = -s
			}
			if s < sub {
				t = total
				sub = s
			}
		}
	}

	return t
}

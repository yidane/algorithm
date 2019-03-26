package problems

import (
	"bytes"
	"container/list"
	"fmt"
	"sort"
	"strconv"
)

/*
 * @lc app=leetcode id=179 lang=golang
 *
 * [179] Largest Number
 *
 * https://leetcode.com/problems/largest-number/description/
 *
 * algorithms
 * Medium (25.34%)
 * Total Accepted:    122.9K
 * Total Submissions: 484.2K
 * Testcase Example:  '[10,2]'
 *
 * Given a list of non negative integers, arrange them such that they form the
 * largest number.
 *
 * Example 1:
 *
 *
 * Input: [10,2]
 * Output: "210"
 *
 * Example 2:
 *
 *
 * Input: [3,30,34,5,9]
 * Output: "9534330"
 *
 *
 * Note: The result may be very large, so you need to return a string instead
 * of an integer.
 *
 */
func largestNumber(nums []int) string {
	if len(nums) == 0 {
		return ""
	}

	if len(nums) == 1 {
		return strconv.Itoa(nums[0])
	}

	quickLargestNumber(nums, 0, len(nums)-1)

	fmt.Println(nums)

	buf := bytes.Buffer{}
	containNoZero := false
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			containNoZero = true
		}

		if containNoZero {
			buf.WriteString(strconv.Itoa(nums[i]))
		}
	}

	if buf.Len() == 0 {
		return "0"
	}

	return buf.String()
}

func quickLargestNumber(nums []int, begin, end int) {
	if begin >= end {
		return
	}

	partition := nums[begin]

	i, j := begin, end
	for i < j {
		for i < j && bigger(partition, nums[j]) {
			j--
		}

		nums[i] = nums[j]

		for i < j && (bigger(nums[i], partition) || nums[i] == partition) {
			i++
		}

		nums[j] = nums[i]
	}

	nums[i] = partition

	quickLargestNumber(nums, begin, i-1)
	quickLargestNumber(nums, i+1, end)
}

//大于
func bigger(a, b int) bool {
	// 39 > 38 > 37 > 36 > 35 > 354 > 3543 > 34 > 3 > 33 > 32 > 31 > 30
	// 12 > 121
	// 122 > 12

	if a == b {
		return false
	}

	as := ints(a)
	bs := ints(b)

	var ai, bi int
	for i := 0; i < len(as)+len(bs); i++ {
		if i < len(as) {
			ai = as[i]
		} else {
			ai = bs[i-len(as)]
		}

		if i < len(bs) {
			bi = bs[i]
		} else {
			bi = as[i-len(bs)]
		}

		if ai == bi {
			continue
		}

		return ai > bi
	}

	return true
}

func ints(i int) []int {
	l := list.New()
	for i >= 10 {
		j := i % 10
		i = i / 10
		l.PushFront(j)
	}
	l.PushFront(i)

	result := make([]int, l.Len())

	head := l.Front()
	n := 0
	for head != nil {
		result[n] = head.Value.(int)
		head = head.Next()
		n++
	}

	return result
}

//best in leetcode
func largestNumber1(nums []int) string {
	arr := make([]string, len(nums))

	for i := 0; i < len(nums); i++ {
		arr[i] = strconv.Itoa(nums[i])
	}

	sort.Slice(arr, func(i, j int) bool {
		as := arr[i]
		bs := arr[j]

		var ai, bi uint8
		for k := 0; k < len(as)+len(bs); k++ {
			if k < len(as) {
				ai = as[k]
			} else {
				ai = bs[k-len(as)]
			}

			if k < len(bs) {
				bi = bs[k]
			} else {
				bi = as[k-len(bs)]
			}

			if ai == bi {
				continue
			}

			return ai > bi
		}

		return true
	})

	buf := bytes.Buffer{}
	containNoZero := false
	for i := 0; i < len(arr); i++ {
		if arr[i] != "0" {
			containNoZero = true
		}

		if containNoZero {
			buf.WriteString(arr[i])
		}
	}

	if buf.Len() == 0 {
		return "0"
	}

	return buf.String()
}

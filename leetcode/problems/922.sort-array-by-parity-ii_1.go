package problems

/*
 * @lc app=leetcode id=922 lang=golang
 *
 * [922] Sort Array By Parity II
 *
 * https://leetcode.com/problems/sort-array-by-parity-ii/description/
 *
 * algorithms
 * Easy (66.73%)
 * Total Accepted:    32.7K
 * Total Submissions: 48.9K
 * Testcase Example:  '[4,2,5,7]'
 *
 * Given an array A of non-negative integers, half of the integers in A are
 * odd, and half of the integers are even.
 *
 * Sort the array so that whenever A[i] is odd, i is odd; and whenever A[i] is
 * even, i is even.
 *
 * You may return any answer array that satisfies this condition.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [4,2,5,7]
 * Output: [4,5,2,7]
 * Explanation: [4,7,2,5], [2,5,4,7], [2,7,4,5] would also have been
 * accepted.
 *
 *
 *
 *
 * Note:
 *
 *
 * 2 <= A.length <= 20000
 * A.length % 2 == 0
 * 0 <= A[i] <= 1000
 *
 *
 *
 *
 *
 */
func sortArrayByParityII(A []int) []int {
	for i, j := 0, 1; i < len(A); i += 2 {
		if A[i]&1 == 0 {
			continue
		}

		//找到奇数位置的第一个偶数
		for ; A[j]&1 == 1; j += 2 {
		}

		//进行交换
		A[i], A[j] = A[j], A[i]
	}

	return A
}

func sortArrayByParityII1(A []int) []int {
	if len(A) == 0 {
		return A
	}

	evenIndex := 0 //偶数
	oddIndex := 1  //奇数

	result := make([]int, len(A))
	for i := 0; i < len(A); i++ {
		b := A[i] % 2

		if b == 0 {
			result[evenIndex] = A[i]
			evenIndex += 2
		} else {
			result[oddIndex] = A[i]
			oddIndex += 2
		}
	}

	return result

}

//使用通道处理
func sortArrayByParityIIChan(A []int) []int {
	result := make([]int, 0)
	evenChan := make(chan int, len(A)/2)
	oddChan := make(chan int, len(A)/2)

	go func() {
		for i := 0; i < len(A); i++ {
			item := A[i]
			if item%2 == 0 {
				evenChan <- item
			} else {
				oddChan <- item
			}
		}
	}()

	for i := 0; i < len(A); i++ {
		if i%2 == 0 {
			result = append(result, <-evenChan)
		} else {
			result = append(result, <-oddChan)
		}
	}

	return result
}

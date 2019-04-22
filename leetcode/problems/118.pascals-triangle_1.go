package problems

/*
 * @lc app=leetcode id=118 lang=golang
 *
 * [118] Pascal's Triangle
 *
 * https://leetcode.com/problems/pascals-triangle/description/
 *
 * algorithms
 * Easy (45.36%)
 * Total Accepted:    243.1K
 * Total Submissions: 535.2K
 * Testcase Example:  '5'
 *
 * Given a non-negative integer numRows, generate the first numRows of Pascal's
 * triangle.
 *
 *
 * In Pascal's triangle, each number is the sum of the two numbers directly
 * above it.
 *
 * Example:
 *
 *
 * Input: 5
 * Output:
 * [
 * ⁠    [1],
 * ⁠   [1,1],
 * ⁠  [1,2,1],
 * ⁠ [1,3,3,1],
 * ⁠[1,4,6,4,1]
 * ]
 *
 *
 */
func generate(numRows int) [][]int {
	if numRows <= 0 {
		return nil
	}

	if numRows == 1 {
		return [][]int{{1}}
	}

	result := [][]int{{1}, {1, 1}}
	for i := 2; i < numRows; i++ {
		lastArr := result[i-1]
		newArr := make([]int, len(lastArr)+1)
		newArr[0] = 1
		for j := 0; j < len(lastArr)-1; j++ {
			newArr[j+1] = lastArr[j] + lastArr[j+1]
		}
		newArr[len(newArr)-1] = 1
		result = append(result, newArr)
	}

	return result
}

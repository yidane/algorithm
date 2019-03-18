package problems

/*
 * @lc app=leetcode id=861 lang=golang
 *
 * [861] Score After Flipping Matrix
 *
 * https://leetcode.com/problems/score-after-flipping-matrix/description/
 *
 * algorithms
 * Medium (68.82%)
 * Total Accepted:    10.3K
 * Total Submissions: 15K
 * Testcase Example:  '[[0,0,1,1],[1,0,1,0],[1,1,0,0]]'
 *
 * We have a two dimensional matrix A where each value is 0 or 1.
 *
 * A move consists of choosing any row or column, and toggling each value in
 * that row or column: changing all 0s to 1s, and all 1s to 0s.
 *
 * After making any number of moves, every row of this matrix is interpreted as
 * a binary number, and the score of the matrix is the sum of these numbers.
 *
 * Return the highest possible score.
 *
 *
 *
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [[0,0,1,1],[1,0,1,0],[1,1,0,0]]
 * Output: 39
 * Explanation:
 * Toggled to [[1,1,1,1],[1,0,0,1],[1,1,1,1]].
 * 0b1111 + 0b1001 + 0b1111 = 15 + 9 + 15 = 39
 *
 *
 *
 * Note:
 *
 *
 * 1 <= A.length <= 20
 * 1 <= A[0].length <= 20
 * A[i][j] is 0 or 1.
 *
 *
 *
 */

/*
 0,0,1,1
 1,0,1,0
 1,1,0,0

 1 1 0 0
 1 0 1 0
 1 1 0 0

 1 1 1 0
 1 0 0 0
 1 1 1 0

 1 1 1 1
 1 0 0 1
 1 1 1 1
*/
func matrixScore(A [][]int) int {
	if len(A) == 0 {
		return 0
	}

	/*
		先行操作，保证第一列都是1
		再按列操作，保证每一列1比0多
	*/

	for i := 0; i < len(A); i++ {
		if A[i][0] != 1 {
			for j := 0; j < len(A[i]); j++ {
				if A[i][j] == 0 {
					A[i][j] = 1
				} else {
					A[i][j] = 0
				}
			}
		}
	}

	for j := 1; j < len(A[0]); j++ {
		count := 0
		for i := 0; i < len(A); i++ {
			if A[i][j] == 1 {
				count++
			}
		}

		if count*2 < len(A) {
			for i := 0; i < len(A); i++ {
				if A[i][j] == 1 {
					A[i][j] = 0
				} else {
					A[i][j] = 1
				}
			}
		}
	}

	total := 0

	for i := 0; i < len(A); i++ {
		var subTotal int
		l := len(A[0])
		for j := 0; j < l; j++ {
			if A[i][j] == 1 {
				subTotal += 1 << uint32(l-j-1)
			}
		}

		total += subTotal
	}

	return total
}

package problems

/*
 * @lc app=leetcode id=365 lang=golang
 *
 * [365] Water and Jug Problem
 *
 * https://leetcode.com/problems/water-and-jug-problem/description/
 *
 * algorithms
 * Medium (28.81%)
 * Total Accepted:    26.7K
 * Total Submissions: 92.6K
 * Testcase Example:  '3\n5\n4'
 *
 * You are given two jugs with capacities x and y litres. There is an infinite
 * amount of water supply available. You need to determine whether it is
 * possible to measure exactly z litres using these two jugs.
 *
 * If z liters of water is measurable, you must have z liters of water
 * contained within one or both buckets by the end.
 *
 * Operations allowed:
 *
 *
 * Fill any of the jugs completely with water.
 * Empty any of the jugs.
 * Pour water from one jug into another till the other jug is completely full
 * or the first jug itself is empty.
 *
 *
 * Example 1: (From the famous "Die Hard" example)
 *
 *
 * Input: x = 3, y = 5, z = 4
 * Output: True
 *
 *
 * Example 2:
 *
 *
 * Input: x = 2, y = 6, z = 5
 * Output: False
 *
 */

/*
z=ax+by
b=(z-ax)/y
*/

//todo: 或许用几何知识能解决
func canMeasureWater(x int, y int, z int) bool {
	if z == 0 {
		return true
	}

	if x < y {
		x, y = y, x
	}

	i := 0
	for {
		temp := z - i*x

		if temp%y == 0 {
			return true
		}

		i++
	}
}

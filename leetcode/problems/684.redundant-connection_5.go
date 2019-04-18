package problems

/*
 * @lc app=leetcode id=684 lang=golang
 *
 * [684] Redundant Connection
 *
 * https://leetcode.com/problems/redundant-connection/description/
 *
 * algorithms
 * Medium (51.10%)
 * Total Accepted:    46.7K
 * Total Submissions: 91.3K
 * Testcase Example:  '[[1,2],[1,3],[2,3]]'
 *
 *
 * In this problem, a tree is an undirected graph that is connected and has no
 * cycles.
 *
 * The given input is a graph that started as a tree with N nodes (with
 * distinct values 1, 2, ..., N), with one additional edge added.  The added
 * edge has two different vertices chosen from 1 to N, and was not an edge that
 * already existed.
 *
 * The resulting graph is given as a 2D-array of edges.  Each element of edges
 * is a pair [u, v] with u < v, that represents an undirected edge connecting
 * nodes u and v.
 *
 * Return an edge that can be removed so that the resulting graph is a tree of
 * N nodes.  If there are multiple answers, return the answer that occurs last
 * in the given 2D-array.  The answer edge [u, v] should be in the same format,
 * with u < v.
 * Example 1:
 *
 * Input: [[1,2], [1,3], [2,3]]
 * Output: [2,3]
 * Explanation: The given undirected graph will be like this:
 * ⁠ 1
 * ⁠/ \
 * 2 - 3
 *
 *
 * Example 2:
 *
 * Input: [[1,2], [2,3], [3,4], [1,4], [1,5]]
 * Output: [1,4]
 * Explanation: The given undirected graph will be like this:
 * 5 - 1 - 2
 * ⁠   |   |
 * ⁠   4 - 3
 *
 *
 * Note:
 * The size of the input 2D-array will be between 3 and 1000.
 * Every integer represented in the 2D-array will be between 1 and N, where N
 * is the size of the input array.
 *
 *
 *
 *
 *
 * Update (2017-09-26):
 * We have overhauled the problem description + test cases and specified
 * clearly the graph is an undirected graph. For the directed graph follow up
 * please see Redundant Connection II). We apologize for any inconvenience
 * caused.
 *
 */

//关键点是判断环存在的位置，若存在环，则解除环

//先构建图，使用DFS可以判断是否存在环，当遍历出现已经遍历过得点，则环出现
//todo:需要好好思考下
func findRedundantConnection(edges [][]int) []int {
	m := map[int]bool{}
	m[edges[0][0]] = true
	m[edges[0][1]] = true
	for i := 1; i < len(edges); i++ {
		edge := edges[i]
		if m[edge[1]] {
			return edge
		}

		m[edge[1]] = true
	}

	return nil
}

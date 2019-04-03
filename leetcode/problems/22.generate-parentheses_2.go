package problems

import "sort"

/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 *
 * https://leetcode.com/problems/generate-parentheses/description/
 *
 * algorithms
 * Medium (53.74%)
 * Total Accepted:    315.4K
 * Total Submissions: 585.9K
 * Testcase Example:  '3'
 *
 *
 * Given n pairs of parentheses, write a function to generate all combinations
 * of well-formed parentheses.
 *
 *
 *
 * For example, given n = 3, a solution set is:
 *
 *
 * [
 * ⁠ "((()))",
 * ⁠ "(()())",
 * ⁠ "(())()",
 * ⁠ "()(())",
 * ⁠ "()()()"
 * ]
 *
 */

/*
1 ()
2 ()() (())
3 ()()() (())() ()(()) ((())) (()())
基本思路是递归实现，每一个表达式有自己的可添加域
*/
func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}

	if n == 1 {
		return []string{"()"}
	}

	if n == 2 {
		return []string{"()()", "(())"}
	}

	resultMap := map[string]bool{"()()": true, "(())": true}
	for i := 2; i < n; i++ {
		resultMap = addParenthesis(resultMap)
	}

	result := make([]string, len(resultMap))

	i := 0
	for r := range resultMap {
		result[i] = r
		i++
	}

	sort.Strings(result)
	return result
}

func addParenthesis(arr map[string]bool) map[string]bool {
	var result = make(map[string]bool)

	//寻找结构（），在结构内添加新的（）即为新的结果
	//也可以寻找（（）），添加为（（）（））
	for cur := range arr {
		j := 0
		for ; j < len(cur)-2; j++ {
			if cur[j] == '(' && cur[j+1] == ')' {
				result[cur[:j+1]+"()"+cur[j+1:]] = true
			}
			if cur[j] == '(' && cur[j+1] == '(' && cur[j+2] == ')' {
				result[cur[:j+1]+"()"+cur[j+1:]] = true
			}
		}

		if cur[j] == '(' && cur[j+1] == ')' {
			result[cur[:j+1]+"()"+cur[j+1:]] = true
		}

		result["()"+cur] = true
		result["("+cur+")"] = true
		result[cur+"()"] = true
	}

	return result
}

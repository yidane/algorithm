package problems

/*
 * @lc app=leetcode id=10 lang=golang
 *
 * [10] Regular Expression Matching
 *
 * https://leetcode.com/problems/regular-expression-matching/description/
 *
 * algorithms
 * Hard (25.00%)
 * Total Accepted:    280.1K
 * Total Submissions: 1.1M
 * Testcase Example:  '"aa"\n"a"'
 *
 * Given an input string (s) and a pattern (p), implement regular expression
 * matching with support for '.' and '*'.
 *
 *
 * '.' Matches any single character.
 * '*' Matches zero or more of the preceding element.
 *
 *
 * The matching should cover the entire input string (not partial).
 *
 * Note:
 *
 *
 * s could be empty and contains only lowercase letters a-z.
 * p could be empty and contains only lowercase letters a-z, and characters
 * like . or *.
 *
 *
 * Example 1:
 *
 *
 * Input:
 * s = "aa"
 * p = "a"
 * Output: false
 * Explanation: "a" does not match the entire string "aa".
 *
 *
 * Example 2:
 *
 *
 * Input:
 * s = "aa"
 * p = "a*"
 * Output: true
 * Explanation: '*' means zero or more of the precedeng element, 'a'.
 * Therefore, by repeating 'a' once, it becomes "aa".
 *
 *
 * Example 3:
 *
 *
 * Input:
 * s = "ab"
 * p = ".*"
 * Output: true
 * Explanation: ".*" means "zero or more (*) of any character (.)".
 *
 *
 * Example 4:
 *
 *
 * Input:
 * s = "aab"
 * p = "c*a*b"
 * Output: true
 * Explanation: c can be repeated 0 times, a can be repeated 1 time. Therefore
 * it matches "aab".
 *
 *
 * Example 5:
 *
 *
 * Input:
 * s = "mississippi"
 * p = "mis*is*p*."
 * Output: false
 *
 *
 */

//todo:bugs
func isMatch(s string, p string) bool {
	var lastRune, nextRune rune = -1, -1

	j := 0
	for _, c := range s {
		if j == len(p) {
			return false
		}

		pc := rune(p[j])

		for {
			if pc == c || pc == '.' {
				//当前字符匹配，游标右移
				lastRune = pc
				if j+1 < len(p) {
					nextRune = rune(p[j+1])
				} else {
					nextRune = -1
				}
				j++
				break
			} else if pc == '*' {
				//当前字符为(*)，判断上一个字符是否匹配
				if lastRune == c || lastRune == '.' {
					if j+1 < len(p) {
						nextRune = rune(p[j+1])
					} else {
						nextRune = -1
					}
					j++
					break
				}
			} else {
				//当前字符不是(.*),判断下一个字符是否存在，且下一个字符为(*),这样可以过滤此匹配
				if j+1 < len(p) {
					np := rune(p[j+1])
					if np != '*' {
						return false
					}
					j++
				}
			}

			j++

			if j == len(p) {
				return false
			}

			pc = rune(p[j])
		}
	}

	return nextRune == -1
}

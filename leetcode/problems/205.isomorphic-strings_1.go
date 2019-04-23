package problems

/*
 * @lc app=leetcode id=205 lang=golang
 *
 * [205] Isomorphic Strings
 *
 * https://leetcode.com/problems/isomorphic-strings/description/
 *
 * algorithms
 * Easy (37.04%)
 * Total Accepted:    196.7K
 * Total Submissions: 530.7K
 * Testcase Example:  '"egg"\n"add"'
 *
 * Given two strings s and t, determine if they are isomorphic.
 *
 * Two strings are isomorphic if the characters in s can be replaced to get t.
 *
 * All occurrences of a character must be replaced with another character while
 * preserving the order of characters. No two characters may map to the same
 * character but a character may map to itself.
 *
 * Example 1:
 *
 *
 * Input: s = "egg", t = "add"
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: s = "foo", t = "bar"
 * Output: false
 *
 * Example 3:
 *
 *
 * Input: s = "paper", t = "title"
 * Output: true
 *
 * Note:
 * You may assume both s and t have the same length.
 *
 */
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	ms := make(map[uint8]uint8, len(s))
	mt := make(map[uint8]uint8, len(s))

	for i := 0; i < len(s); i++ {
		si := s[i]
		ti := t[i]

		if v, ok := ms[si]; ok {
			if v != ti {
				return false
			}
		} else {
			ms[si] = ti
		}

		if v, ok := mt[ti]; ok {
			if v != si {
				return false
			}
		} else {
			mt[ti] = si
		}
	}

	return true
}

package problems

/*
 * @lc app=leetcode id=30 lang=golang
 *
 * [30] Substring with Concatenation of All Words
 *
 * https://leetcode.com/problems/substring-with-concatenation-of-all-words/description/
 *
 * algorithms
 * Hard (23.32%)
 * Total Accepted:    127K
 * Total Submissions: 544.3K
 * Testcase Example:  '"barfoothefoobarman"\n["foo","bar"]'
 *
 * You are given a string, s, and a list of words, words, that are all of the
 * same length. Find all starting indices of substring(s) in s that is a
 * concatenation of each word in words exactly once and without any intervening
 * characters.
 *
 * Example 1:
 *
 *
 * Input:
 * ⁠ s = "barfoothefoobarman",
 * ⁠ words = ["foo","bar"]
 * Output: [0,9]
 * Explanation: Substrings starting at index 0 and 9 are "barfoor" and "foobar"
 * respectively.
 * The output order does not matter, returning [9,0] is fine too.
 *
 *
 * Example 2:
 *
 *
 * Input:
 * ⁠ s = "wordgoodgoodgoodbestword",
 * ⁠ words = ["word","good","best","word"]
 * Output: []
 *
 *
 */
func findSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return []int{}
	}

	itemLength := len(words[0])
	totalLength := len(words) * itemLength
	if totalLength > len(s) {
		return []int{}
	}

	wordsMap := make(map[string][]int, len(words))
	for i := 0; i < len(words); i++ {
		wordsMap[words[i]] = append(wordsMap[words[i]], i+1)
	}
	total := (1 + len(words)) * len(words) / 2

	var result []int

	for i := 0; i <= len(s)-totalLength; i++ {
		s1 := s[i : i+totalLength]
		if matchSubstring(s1, wordsMap, itemLength, total) {
			result = append(result, i)
		}
	}

	return result
}

//words可能会出现重复项
func matchSubstring(s string, words map[string][]int, l, t int) bool {
	m := make(map[int]bool)

	for i := 0; i < len(s); i += l {
		s1 := s[i : i+l]
		arr := words[s1]

		if len(arr) == 0 {
			return false
		}

		for j := 0; j < len(arr); j++ {
			if !m[arr[j]] {
				m[arr[j]] = true
				t = t - arr[j]
				break
			}
		}
	}

	return t == 0
}

package problems

/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 *
 * https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
 *
 * algorithms
 * Medium (28.05%)
 * Total Accepted:    841.5K
 * Total Submissions: 3M
 * Testcase Example:  '"abcabcbb"'
 *
 * Given a string, find the length of the longest substring without repeating
 * characters.
 *
 *
 * Example 1:
 *
 *
 * Input: "abcabcbb"
 * Output: 3
 * Explanation: The answer is "abc", with the length of 3.
 *
 *
 *
 * Example 2:
 *
 *
 * Input: "bbbbb"
 * Output: 1
 * Explanation: The answer is "b", with the length of 1.
 *
 *
 *
 * Example 3:
 *
 *
 * Input: "pwwkew"
 * Output: 3
 * Explanation: The answer is "wke", with the length of 3.
 * ⁠            Note that the answer must be a substring, "pwke" is a
 * subsequence and not a substring.
 *
 *
 *
 *
 */
/*
每次都判断当前字符是否已经出现，若出现，则当前位置和起始位置的差即为当前最长子串。
然后，将开始位置移到当前位置移到当前字符上一次出现重复的后一位。
	为什么不对开始位置顺延，因为后续已经判断出有重复子串为当前字符的重复
	若仅仅是顺延，很明显比不顺延要小一位，不是最长子串，
	若移到当前字符上次出现位置的下一位，这样就排除了当前字符的重复，
	可以对后续字符继续判断。
如此遍历完毕之后，还需要判断最后一段字符串长度是否大于上次得出的最长子串。

算法关键是对开始位置的移位，移位到当前字符上次出现的位置的下一位。
*/
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	bArr := make(map[rune]int)
	maxLength := 0
	begin := 0
	for i, v := range s {
		if j, ok := bArr[v]; ok {
			if i-begin > maxLength {
				maxLength = i - begin
			}
			if j+1 > begin {
				begin = j + 1
			}
		}
		bArr[v] = i
	}

	if len(s)-begin > maxLength {
		maxLength = len(s) - begin
	}

	return maxLength
}

package problems

/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 *
 * https://leetcode.com/problems/longest-palindromic-substring/description/
 *
 * algorithms
 * Medium (26.85%)
 * Total Accepted:    509.5K
 * Total Submissions: 1.9M
 * Testcase Example:  '"babad"'
 *
 * Given a string s, find the longest palindromic substring in s. You may
 * assume that the maximum length of s is 1000.
 *
 * Example 1:
 *
 *
 * Input: "babad"
 * Output: "bab"
 * Note: "aba" is also a valid answer.
 *
 *
 * Example 2:
 *
 *
 * Input: "cbbd"
 * Output: "bb"
 *
 *
 */

/*
以下算法更好：马拉车算法
http://www.cnblogs.com/grandyang/p/4475985.html
*/
func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	m := map[uint8][]int{}
	for i := 0; i < len(s); i++ {
		arr := m[s[i]]
		arr = append(arr, i)
		m[s[i]] = arr
	}

	if len(m) == len(s) {
		return s[:1]
	}

	begin, end := 0, 1
	for _, arr := range m {
		if len(arr) == 1 {
			continue
		}

		//i,j之间的数据，若是偶数，肯定是成对出现。若是奇数，则最中间之外的数据是成对出现
		for i, j := 0, len(arr)-1; i < j; {
			head := arr[i]
			tail := arr[j] + 1

			if tail-head <= end-begin {
				break
			}

			if tail-head == 2 {
				begin = head
				end = tail
				break
			}

			for h := i; h < j; h++ {
				head1 := arr[h]
				if tail-head1 <= end-begin {
					break
				}

				if isPalindromeString(s[head1:tail]) {
					begin = head1
					end = tail
					break
				}
			}

			for t := j; t > i; t-- {
				tail1 := arr[t] + 1
				if tail1-head <= end-begin {
					break
				}

				if isPalindromeString(s[head:tail1]) {
					begin = head
					end = tail1
					break
				}
			}

			i++
			j--
		}
	}

	return s[begin:end]
}

func isPalindromeString(s string) bool {
	if len(s) < 2 {
		return true
	}

	for i, j := 0, len(s)-1; i < j; {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}
